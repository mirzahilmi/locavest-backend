package cfg

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func NewDatabase() *sqlx.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	schema := os.Getenv("DB_SCHEMA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Caller().Msg("")
	}

	return db
}
