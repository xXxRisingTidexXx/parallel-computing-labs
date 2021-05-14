package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
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
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	<-signals
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
