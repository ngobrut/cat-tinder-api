package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ngobrut/cat-tinder-api/config"
)

func NewDBClient(cnf *config.Config) (*sql.DB, error) {
	host := cnf.Postgres.Host
	port := cnf.Postgres.Port
	user := cnf.Postgres.User
	password := cnf.Postgres.Password
	dbname := cnf.Postgres.Database
	params := cnf.Postgres.Params

	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, dbname, params)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Println("[failed-postgres-connection]", err)
		return nil, err
	}

	log.Println("connected to postgres")

	return db, nil
}
