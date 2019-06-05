package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	lines, err := linesCount(logPath)
	if err != nil {
		log.Fatalf("Error counting lines in log file of %s at %s: %s", n, d, err)
	}

	err = linesPrint(logPath, p, lines, r)
	if err != nil {
		log.Fatalf("Error printing log file of %s at %s: %s", n, d, err)
	}
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

func linesPrint(filePath, period string, count, rows int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	timeForPeriod := time.Now()
	timeForPeriodStr := timeForPeriod.Format("Jan  2 15:04:05")
	timeForPeriod, err = time.Parse("Jan  2 15:04:05", timeForPeriodStr)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	i := 0
	if rows != -1 {
		for scanner.Scan() {
			i++
			if i > count-rows {
				line := scanner.Text()

				fmt.Printf("%d: %s\n", i, line)
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	} else {
		unit := period[strings.LastIndexAny(period, "1234567890")+1 : len(period)]
		num := period[0:strings.Index(period, unit)]
		numFloat, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return err
		}

		for scanner.Scan() {
			i++

			line := scanner.Text()
			timeOfLogStr := line[0 : strings.Index(line, strings.Split(filePath, "/")[4])-1]
			timeOfLog, err := time.Parse("Jan  2 15:04:05", timeOfLogStr)
			if err != nil {
				return err
			}

			duration := timeForPeriod.Sub(timeOfLog)
			switch unit {
			case "h", "hours":
				{
					if duration.Hours() < numFloat {
						fmt.Printf("%d: %s\n", i, line)
					}
				}
			case "m", "minutes":
				{
					if duration.Minutes() < numFloat {
						fmt.Printf("%d: %s\n", i, line)
					}
				}
			case "s", "seconds":
				{
					if duration.Seconds() < numFloat {
						fmt.Printf("%d: %s\n", i, line)
					}
				}
			default:
				{
					return errors.New("Unknow time unit")
				}
			}
		}
	}

	return nil
}
