package main

import (
	"errors"
	"net/http"

	"github.com/go-ldap/ldap"
	_ "github.com/kshvakov/clickhouse"
	log "github.com/sirupsen/logrus"
	"github.com/ultram4rine/logviewer/db"
	"github.com/ultram4rine/logviewer/server"
)

func main() {
	var (
		port     = ":4027"
		confPath = "conf.json"
	)

	err := server.Init(confPath)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLogin(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		date := r.FormValue("date")
		name := r.FormValue("name")
		time := r.FormValue("time")

		logs, err := db.GetLogfromSwitch(name, time)
		if err != nil {
			log.Printf("Error printing log file of %s at %s: %s", name, date, err)
		}

		w.Write([]byte(logs))
	})

	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLogin(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		http.ServeFile(w, r, "public/html/index.html")
	})

	log.Println("Starting...")
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func auth(login, password string) (string, error) {
	if password == "" {
		return "", errors.New("Empty password")
	}

	username := ""

	l, err := ldap.Dial("tcp", server.Config.LdapServer)
	if err != nil {
		return username, err
	}
	defer l.Close()

	if l.Bind(server.Config.LdapUser, server.Config.LdapPassword); err != nil {
		return username, err
	}

	searchRequest := ldap.NewSearchRequest(
		server.Config.LdapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(sAMAccountName="+login+"))",
		[]string{"cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil || len(sr.Entries) != 1 {
		return username, errors.New("User not found")
	}

	username = sr.Entries[0].GetAttributeValue("cn")

	if err = l.Bind(username, password); err != nil {
		return "", err
	}

	return username, err
}

func alreadyLogin(r *http.Request) bool {
	session, _ := server.Server.Store.Get(r, "logviewer_session")
	return session.Values["userName"] != nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := server.Server.Store.Get(r, "logviewer_session")

	if r.Method == "GET" {
		http.ServeFile(w, r, "public/html/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		if alreadyLogin(r) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		userName, err := auth(r.FormValue("uname"), r.FormValue("psw"))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		session.Values["userName"] = userName
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
