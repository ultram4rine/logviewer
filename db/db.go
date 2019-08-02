package db

import (
	"fmt"
	"time"

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
	LogTime      time.Time `db:"log_time"`
	LogEventNum  string    `db:"log_event_number"`
	LogModule    string    `db:"log_module"`
	LogMessage   string    `db:"log_msg"`
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

func GetLogfromSwitch(swName, period string) (string, error) {
	var (
		ls   []switchLog
		logs string
	)

	if err := server.Server.DB.Select(&ls, "SELECT * FROM switchlogs WHERE sw_name = ?", swName); err != nil {
		return "", err
	}

	for i, l := range ls {
		logs += fmt.Sprintf("%d: %s\n", i, l.LogMessage)
	}

	return logs, nil
}
