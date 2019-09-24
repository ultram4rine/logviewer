package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-ldap/ldap"
	log "github.com/sirupsen/logrus"

	"github.com/ultram4rine/logviewer/db"
	"github.com/ultram4rine/logviewer/server"
)

//RootHandler handle root path
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	http.ServeFile(w, r, "public/html/index.html")
}

//GetHandler handle get request
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}

	getType := r.FormValue("type")

	switch getType {
	case "sw":
		{
			name := r.FormValue("name")
			time := r.FormValue("time")
			periodInt, err := strconv.Atoi(time)
			if err != nil {
				log.Warnf("Error parsing time: %v", err)
			}

			logs, err := db.GetLogfromSwitch(name, periodInt)
			if err != nil {
				log.Warnf("Error printing log file of %s: %v", name, err)
			}

			w.Write([]byte(logs))
		}
	case "dhcp":
		{
			mac := r.FormValue("mac")

			result, err := db.GetDHCPLogs(mac)
			if err != nil {
				log.Warnf("Error with geting dhcp logs: %v", err)
			}

			logsJSON, err := json.Marshal(result)
			if err != nil {
				log.Warn(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(logsJSON)
		}
	}
}

//LoginHandler handle login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
