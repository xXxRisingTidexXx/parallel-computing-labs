package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("SEARCH_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
