package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/ultram4rine/logviewer/helpers"
)

func main() {
	var port = ":4027"

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		date := r.FormValue("date")
		ip := r.FormValue("ip")
		time := r.FormValue("time")

		logPath := "/var/log/remote/" + ip + "/" + date

		lines, err := helpers.LinesCount(logPath)
		if err != nil {
			log.Printf("Error counting lines in log file of %s at %s: %s", ip, date, err)
		}

		logs, err := helpers.Lines2String(logPath, time, lines, -1)
		if err != nil {
			log.Printf("Error printing log file of %s at %s: %s", ip, date, err)
		}

		w.Write([]byte(logs))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/html/index.html")
	})

	log.Println("Starting...")
	err := http.ListenAndServe("localhost"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
