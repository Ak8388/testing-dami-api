package database

import (
	"dami-api/app/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	cfg := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME, cfg.DBPORT)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	db.Ping()

	return db
}
