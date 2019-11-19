package db

import (
	"net"
	"time"

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

type dhcpLog struct {
	Timestamp     time.Time `db:"ts"`
	TimeStampStr  string
	TimestampNano int64  `db:"ts_nano"`
	Message       string `db:"message"`
	Server        string `db:"server"`
	Severity      string `db:"severity"`
	IP            net.IP `db:"ip"`
	MAC           uint64 `db:"mac"`
	MACStr        string
	Request       string `db:"request"`
	ServerID      string `db:"server_id"`
	ClientHost    string `db:"client_host"`
	Link          string `db:"link"`
	Extra         string `db:"extra"`
	ReverseName   string `db:"reverse_name"`
	DNSName       string `db:"dns_name"`
	Reason        string `db:"reason"`
	Subnet        string `db:"subnet"`
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
	var logs []switchLog

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

func GetDHCPLogs(mac string, period int) ([]dhcpLog, error) {
	var logs []dhcpLog

	duration := time.Minute * -time.Duration(period)

	time := time.Now().Add(duration)

	if err := server.Server.DB.Select(&logs, "SELECT ts, message, ip FROM dhcp.events WHERE mac = MACStringToNum(?) AND ts > ? ORDER BY ts DESC", mac, time); err != nil {
		return nil, err
	}

	for i := range logs {
		logs[i].MACStr = mac
		logs[i].TimeStampStr = logs[i].Timestamp.Format("2006-01-02 15:04:05")
	}

	return logs, nil
}
