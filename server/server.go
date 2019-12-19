package server

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

var Config struct {
	ListenHTTP    []string `envconfig:"LISTEN_HTTP"`
	LdapUser      string   `envconfig:"LDAP_USER"`
	LdapPassword  string   `envconfig:"LDAP_PASSWORD"`
	LdapServer    string   `envconfig:"LDAP_SERVER"`
	LdapBaseDN    string   `envconfig:"LDAP_BASE_DN"`
	DBHost        string   `envconfig:"DB_HOST" default:"localhost"`
	DBName        string   `envconfig:"DB_NAME" default:"logs"`
	DBUser        string   `envconfig:"DB_USER" default:"root"`
	DBPassword    string   `envconfig:"DB_PASSWORD"`
	SessionKey    string   `envconfig:"SESSION_KEY"`
	EncryptionKey string   `envconfig:"ENCRYPTION_KEY"`
}

var Server struct {
	DB    *sqlx.DB
	Store *sessions.CookieStore
}

func Init() (err error) {
	if err = envconfig.Process("hosting", &Config); err != nil {
		return
	}

	if Config.SessionKey == "" {
		err = errors.New("Empty session key")
		return
	}

	Server.Store = sessions.NewCookieStore([]byte(Config.SessionKey), []byte(Config.EncryptionKey))

	Server.DB, err = sqlx.Open("clickhouse", Config.DBHost+"?username="+Config.DBUser+"&password="+Config.DBPassword+"&database="+Config.DBName)
	if err != nil {
		return
	}

	return nil
}
