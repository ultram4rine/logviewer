package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

var (
	t    = time.Now()
	name = kingpin.Flag("switch", "Switch which log need view").Default().String()
	date = kingpin.Flag("date", "Date of logs").Default(t.Format("2006-01-02")).String()
)

func main() {
	kingpin.Parse()

	var (
		n = *name
		d = *date
	)

	logFile, err := os.Open("/var/log/remote/" + n + "/" + d)
	if err != nil {
		log.Printf("Error opening log file for %s at %s: %s", n, d, err)
	}
	defer logFile.Close()
}
