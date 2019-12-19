package main

import (
	"net/http"

	"github.com/ultram4rine/logviewer/handlers"
	"github.com/ultram4rine/logviewer/server"

	"git.sgu.ru/sgu/systemdutil"
	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/gorilla/mux"
	_ "github.com/kshvakov/clickhouse"
	log "github.com/sirupsen/logrus"
)

func main() {
	systemdutil.Init()

	err := server.Init()
	if err != nil {
		log.Fatalf("Can't init programm: %v", err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	router.HandleFunc("/get/switch", handlers.GetSwitchLogsHandler).Methods("GET")
	router.HandleFunc("/get/dhcp", handlers.GetDHCPLogsHandler).Methods("GET")

	router.HandleFunc("/getavailable", handlers.SimilarHandler).Methods("GET")
	router.HandleFunc("/findsimilar", handlers.SimilarHandler).Methods("GET")

	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/", handlers.RootHandler)

	_, httpSockets, err := systemdutil.ListenSystemdEx(systemdutil.ActivationFiles())
	if err != nil {
		log.Fatal(err)
	}
	httpSockets = append(httpSockets, systemdutil.MustListenTCPSlice(server.Config.ListenHTTP)...)
	systemdutil.ServeAll(nil, httpSockets, nil)

	daemon.SdNotify(false, daemon.SdNotifyReady)
	systemdutil.WaitSigint()
	daemon.SdNotify(false, daemon.SdNotifyStopping)
}
