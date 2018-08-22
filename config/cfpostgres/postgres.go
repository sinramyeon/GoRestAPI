package cfpostgres

import (
	"log"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

// Postgres...
type Postgres struct {
	DB *pgx.Conn

	Host     string
	User     string
	Password string
	Database string
}

// configDB...
// config Postgres
func configDB(p Postgres) pgx.ConnConfig {

	var config pgx.ConnConfig
	config.Host = p.Host
	config.User = p.User
	config.Password = p.Password
	config.Database = p.Database

	return config

}

// NewPostgres...
// Make New Postgresql connection
func NewPostgres(p Postgres) (*pgx.Conn, error) {

	dbconf := configDB(p)

	Postgres, err := pgx.Connect(dbconf)

	if err != nil {
		log.Fatalln("[ERROR] postgres.NewPostgres() '%s'", err)
	}

	return Postgres, err
}
