package main

import (
	"database/sql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("PP_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
	queries := [][]string{
		{"hello"},
		{"non-existing-term"},
		{"non-existing-term"},
		{"qwkd9", "UYB2EBD", "A"},
		{"Com", "n", "give", "w"},
	}
	start := time.Now()
	for _, terms := range queries {
		if _, err := readDocuments(db, terms); err != nil {
			_ = db.Close()
			log.Fatal(err)
		}
	}
	sequentialLatency := time.Now().Sub(start).Microseconds()
	var group sync.WaitGroup
	group.Add(len(queries))
	start = time.Now()
	for _, terms := range queries {
		go readDocumentsWithLog(db, terms, &group)
	}
	group.Wait()
	parallelLatency := time.Now().Sub(start).Microseconds()
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	log.Infof("Sequantial: %d us", sequentialLatency)
	log.Infof("Parallel: %d us", parallelLatency)
}

func readDocuments(db *sql.DB, terms []string) ([]string, error) {
	rows, err := db.Query(
		`select text
		from documents
		where id in (
		    select document_id
		    from entries
		        join tokens on tokens.id = entries.token_id
		    where tokens.text = any($1)
		)`,
		pq.Array(terms),
	)
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

func readDocumentsWithLog(db *sql.DB, terms []string, group *sync.WaitGroup) {
	if _, err := readDocuments(db, terms); err != nil {
		log.Error(err)
	}
	group.Done()
}
