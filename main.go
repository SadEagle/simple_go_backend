package main

import (
	"log"
	"movie_backend_go/db"
)

var c = db.Config{
	Host:     "localhost",
	Port:     5432,
	Database: "movie_server",
	User:     "movie_manager",
	Password: "123",
	SSLMode:  "disable",
}

func main() {
	db, err := db.InitDB(c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
