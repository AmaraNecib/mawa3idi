package main

import (
	"log"
	"mawa3id/DB"
	database "mawa3id/DataBase"
	"mawa3id/api"

	_ "github.com/lib/pq"
)

var DATABASE, err = database.ConnectToDB()

func main() {

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		database.CloseDB(DATABASE)
	}
	queries := DB.New(DATABASE)
	_, err = api.Init(queries)
	if err != nil {
		panic(err)
	}
	defer DATABASE.Close()
}
