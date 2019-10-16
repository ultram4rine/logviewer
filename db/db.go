package db

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/ultram4rine/logviewer/server"
)

type switchLog struct {
	TimeLocal       time.Time `db:"ts_local"`
	SwName          string    `db:"sw_name"`
	SwIP            net.IP    `db:"sw_ip"`
	LogTimeStamp    time.Time `db:"ts_remote"`
	LogTimeStampStr string
	LogFacility     uint8  `db:"facility"`
	LogSeverity     uint8  `db:"severity"`
	LogPriority     uint8  `db:"priority"`
	LogMessage      string `db:"log_msg"`
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

func GetAvailableSwitches() ([]switchLog, error) {
	var switches []switchLog

	rows, err := server.Server.DB.Query("SELECT DISTINCT sw_name, sw_ip FROM switchlogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s switchLog

		if err = rows.Scan(&s.SwName, &s.SwIP); err != nil {
			return nil, err
		}

		switches = append(switches, s)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return switches, nil
}

func GetSimilarSwitches(t string) ([]switchLog, error) {
	var switches []switchLog

	rows, err := server.Server.DB.Query("SELECT DISTINCT sw_name, sw_ip FROM switchlogs WHERE sw_name LIKE ?", t+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s switchLog

		if err = rows.Scan(&s.SwName, &s.SwIP); err != nil {
			return nil, err
		}

		switches = append(switches, s)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return switches, nil
}

func GetLogfromSwitch(swName string, period int) ([]switchLog, error) {
	var (
		logs []switchLog
	)

	duration := time.Minute * -time.Duration(period)

	time := time.Now().Add(duration)

	if err := server.Server.DB.Select(&logs, "SELECT ts_remote, log_msg FROM switchlogs WHERE sw_name = ? AND ts_local > ? ORDER BY ts_local DESC", swName, time); err != nil {
		return nil, err
	}

	for i := range logs {
		logs[i].LogTimeStampStr = logs[i].LogTimeStamp.Format("2006-01-02 15:04:05")
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

			item.Time = time.Unix(0, item.Timestamp*int64(time.Millisecond)).Format("02-Jan-2006 15:04:05")
			result = append(result, item)
		}
	}

	return result, nil
}
