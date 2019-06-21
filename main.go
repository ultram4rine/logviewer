package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/go-ldap/ldap"
	"github.com/gorilla/sessions"
	"github.com/ultram4rine/logviewer/helpers"
)

var store *sessions.CookieStore

var config struct {
	LdapUser      string `json:"ldapUser"`
	LdapPassword  string `json:"ldapPassword"`
	LdapServer    string `json:"ldapServer"`
	LdapBaseDN    string `json:"ldapBaseDN"`
	SessionKey    string `json:"sessionKey"`
}

func main() {
	var (
		port     = ":4027"
		confPath = "conf.json"
	)

	confdata, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(confdata, &config)
	if err != nil {
		log.Fatal(err)
	}
	
	if config.SessionKey == "" {
		log.Fatal(errors.New("Empty session key"))
	}
	store = sessions.NewCookieStore([]byte(config.SessionKey))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLogin(r){
			http.Redirect(w, r, "/login", 301)
		}

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

	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLogin(r){
			http.Redirect(w, r, "/login", 301)
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

	l, err := ldap.Dial("tcp", config.LdapServer)
	if err != nil {
		return username, err
	}
	defer l.Close()

	if l.Bind(config.LdapUser, config.LdapPassword); err != nil {
		return username, err
	}

	searchRequest := ldap.NewSearchRequest(
		config.LdapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(sAMAccountName="+login+"))",
		[]string{"cn"},
		nil,
	)

	if sr, err := l.Search(searchRequest); err != nil || len(sr.Entries) != 1 {
		return username, errors.New("User not found")
	} else {
		username = sr.Entries[0].GetAttributeValue("cn")
	}

	if err = l.Bind(username, password); err != nil {
		return "", err
	}

	return username, err
}

func alreadyLogin(r *http.Request) bool {
	session, _ := store.Get(r, "logviewer_session")
	return session.Values["userName"] != nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "logviewer_session")

	if r.Method == "GET" {
		http.ServeFile(w, r, "public/html/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		if alreadyLogin(r) {
			http.Redirect(w, r, "/", 301)
			return
		}

		if userName, err := auth(r.FormValue("uname"), r.FormValue("psw")); err != nil {
			http.Redirect(w, r, "/login", 301)
			return
		} else {
			session.Values["userName"] = userName
			session.Save(r, w)
			http.Redirect(w, r, "/", 301)
		}
	}
}