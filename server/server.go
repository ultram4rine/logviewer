package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

var Config struct {
	LdapUser      string `json:"ldapUser"`
	LdapPassword  string `json:"ldapPassword"`
	LdapServer    string `json:"ldapServer"`
	LdapBaseDN    string `json:"ldapBaseDN"`
	DBHost        string `json:"dbHost"`
	DBName        string `json:"dbName"`
	DBUser        string `json:"dbUser"`
	DBPassword    string `json:"dbPassword"`
	SessionKey    string `json:"sessionKey"`
	EncryptionKey string `json:"encryptionKey"`
	ElasticServer string `json:"ns.sgu.ru:9200"`
}

var Server struct {
	DB    *sqlx.DB
	Store *sessions.CookieStore
}

func Init(confPath string) error {
	confdata, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(confdata, &Config)
	if err != nil {
		return err
	}

	if Config.SessionKey == "" {
		return errors.New("Empty session key")
	}

	Server.Store = sessions.NewCookieStore([]byte(Config.SessionKey), []byte(Config.EncryptionKey))
	Server.Store.MaxAge(3600)

	Server.DB, err = sqlx.Open("clickhouse", Config.DBHost+"?username="+Config.DBUser+"&password="+Config.DBPassword+"&database="+Config.DBName)
	if err != nil {
		return err
	}

	return nil
}
