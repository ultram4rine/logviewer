package main

import (
	"bufio"
	"fmt"
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
	name = kingpin.Flag("switch", "Switch which log need view").Short('s').Default().Required().String()
	date = kingpin.Flag("date", "Date of logs").Short('d').Default(t.Format("2006-01-02")).String()
	rows = kingpin.Flag("rows", "Number of rows to view").Short('r').Default("40").Int()
)

func main() {
	kingpin.Parse()

	var (
		n = *name
		d = *date
		//r = *rows
		logPath = "/var/log/remote/" + n + "/" + d
	)

	lines, err := linesCount(logPath)
	if err != nil {
		log.Fatalf("Error counting lines in log file of %s at %s: %s", n, d, err)
	}
	fmt.Println(lines)
}

func linesCount(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return count, nil
}
