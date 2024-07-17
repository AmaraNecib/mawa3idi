package main

import (
	"database/sql"
	"log"
	"mawa3id/DB"
	"mawa3id/api"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serviceURI := os.Getenv("DATABASE_URL")

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err := sql.Open("postgres", conn.String())

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := DB.New(db)
	_, err = api.Init(queries)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
