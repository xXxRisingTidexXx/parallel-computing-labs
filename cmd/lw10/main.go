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
	documents, err := readDocuments(db)
	if  err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	log.Infof("Read %d documents", len(documents))
}

func readDocuments(db *sql.DB, terms ...string) ([]string, error) {
	rows, err := db.Query(`select text from documents`)
	if err != nil {
		return nil, err
	}
	documents := make([]string, 0)
	for rows.Next() {
		var document string
		if err := rows.Scan(&document); err != nil {
			_ = rows.Close()
			return nil, err
		}
		documents = append(documents, document)
	}
	if err := rows.Err(); err != nil {
		_ = rows.Close()
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return documents, nil
}
