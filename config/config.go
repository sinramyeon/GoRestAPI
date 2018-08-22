package config

import (
	"GoRestAPI/config/cfpostgres"
	"GoRestAPI/util"

	"log"
	"os"
	"strings"

	"github.com/jackc/pgx"
)

// Configuration ...
type Configuration struct {
	Postgres cfpostgres.Postgres
}

// Global variable
var (

	// ENV is Configuration instance.
	ENV *Configuration
	SQL *pgx.Conn
)

// Init() ...
// Init Config Files and Setup DB
func Init() {
	err := newConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	SQL, err = cfpostgres.NewPostgres(ENV.Postgres)
	if err != nil {
		log.Fatalln(err)
	}
	SQL.Exec(InitSQL)

}

// newConfiguration ...
// make new configuration by config.json file
func newConfiguration() error {
	ENV = &Configuration{}

	var path string

	pwd, _ := os.Getwd()
	if strings.Contains(pwd, "config") {
		path = pwd + "/config/config.json"

	} else {
		os.Chdir("../config")
		pwd, _ = os.Getwd()
		path = pwd + "/config.json"
	}

	if !strings.Contains(path, "config/") {
		path = pwd + "/config/config.json"
	}
	return util.RequireJSON(path, ENV)
}
