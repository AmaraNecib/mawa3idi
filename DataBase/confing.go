package database

import (
	"database/sql"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	serviceURI := os.Getenv("DATABASE_URL")

	conn, _ := url.Parse(serviceURI)

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
