package main

import (
	"time"

	"github.com/ultram4rine/logviewer/helpers"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	t      = time.Now()
	name   = kingpin.Flag("switch", "Switch which log need view").Short('s').Default().Required().String()
	date   = kingpin.Flag("date", "Date of logs").Short('d').Default(t.Format("2006-01-02")).String()
	rows   = kingpin.Flag("rows", "Number of rows to view").Short('r').Default("-1").Int()
	period = kingpin.Flag("period", "Period of last logs").Short('p').Default("30m").String()
)

func main() {
	kingpin.Parse()

	var (
		n       = *name
		d       = *date
		r       = *rows
		p       = *period
		logPath = "/var/log/remote/" + n + "/" + d
	)

	lines, err := helpers.LinesCount(logPath)
	if err != nil {
		log.Fatalf("Error counting lines in log file of %s at %s: %s", n, d, err)
	}

	err = helpers.LinesPrint(logPath, p, lines, r)
	if err != nil {
		log.Fatalf("Error printing log file of %s at %s: %s", n, d, err)
	}
}
