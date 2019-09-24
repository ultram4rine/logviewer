package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/ultram4rine/logviewer/server"
)

type switchLog struct {
	TimeLocal    time.Time `db:"ts_local"`
	SwName       string    `db:"sw_name"`
	SwIP         string    `db:"sw_ip"`
	LogTimeStamp time.Time `db:"ts_remote"`
	LogFacility  int       `db:"facility"`
	LogSeverity  int       `db:"severity"`
	LogPriority  int       `db:"priority"`
	LogTime      string    `db:"log_time"`
	LogEventNum  string    `db:"log_event_number"`
	LogModule    string    `db:"log_module"`
	LogMessage   string    `db:"log_msg"`
}

type LogEntry struct {
	Mac       string
	IP        string
	Timestamp int64
	Time      string
	Link      string
	Message   string
	Request   string
}

func GetAvailableSwitches() (map[string]string, error) {
	var switches = make(map[string]string)

	rows, err := server.Server.DB.Query("SELECT sw_name, sw_ip FROM switchlogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name string
			ip   string
		)

		if err = rows.Scan(&name, &ip); err != nil {
			return nil, err
		}

		switches[name] = ip
	}
	if rows.Err() != nil {
		return nil, err
	}

	return switches, nil
}

func GetLogfromSwitch(swName string, period int) (string, error) {
	var (
		ls   []switchLog
		logs string
	)

	duration := time.Minute * -time.Duration(period)

	time := time.Now().Add(duration)

	if err := server.Server.DB.Select(&ls, "SELECT * FROM switchlogs WHERE sw_name = ? AND ts_remote > ? ORDER BY ts_local", swName, time); err != nil {
		return "", err
	}

	for _, l := range ls {
		logs += fmt.Sprintf("%s: %s\n", l.LogTime, l.LogMessage)
	}

	return logs, nil
}

//GetDHCPLogs geting logs from elasticSearch
func GetDHCPLogs(mac string) ([]LogEntry, error) {
	termQuery := elastic.NewTermQuery("mac", mac)
	searchResult, err := server.Server.ElasticClient.Search().
		Index("dhcp").
		Query(termQuery).
		Sort("timestamp", false).
		From(0).Size(20).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	var result []LogEntry
	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var item LogEntry
			err := json.Unmarshal(hit.Source, &item)
			if err != nil {
				return result, err
			}

			item.Time = time.Unix(0, item.Timestamp*int64(time.Millisecond)).Format("15:04:05 02-Jan-2006")
			result = append(result, item)
		}
	}

	return result, nil
}
