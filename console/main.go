package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
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
