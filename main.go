package main

import (
	"flag"
	"net/http"

	_ "github.com/kshvakov/clickhouse"
	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/logviewer/handlers"
	"github.com/ultram4rine/logviewer/server"
)

var configPath = flag.String("c", "conf.json", "Path to logviewer config json")

func main() {
	flag.Parse()

	err := server.Init(*configPath)
	if err != nil {
		log.Fatalf("Can't init programm: %v", err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/get", handlers.GetHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/", handlers.RootHandler)

	log.Println("Starting server on " + server.Config.Port + " port")
	err = http.ListenAndServe(server.Config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
