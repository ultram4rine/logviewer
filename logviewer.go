package main

import (
	"flag"
	"net/http"

	"github.com/ultram4rine/logviewer/handlers"
	"github.com/ultram4rine/logviewer/server"

	"github.com/gorilla/mux"
	_ "github.com/kshvakov/clickhouse"
	log "github.com/sirupsen/logrus"
)

var configPath = flag.String("c", "logviewer.conf.json", "Path to logviewer config json")

func main() {
	flag.Parse()

	err := server.Init(*configPath)
	if err != nil {
		log.Fatalf("Can't init programm: %v", err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	router.HandleFunc("/get", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/getavailable", handlers.SimilarHandler).Methods("GET")
	router.HandleFunc("/findsimilar", handlers.SimilarHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/", handlers.RootHandler)

	log.Println("Starting server on " + server.Config.Port + " port")
	log.Fatal(http.ListenAndServe(server.Config.Port, router))
}
