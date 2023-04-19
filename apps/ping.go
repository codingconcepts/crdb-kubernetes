package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.SetFlags(0)

	db, err := pgxpool.New(context.Background(), "postgres://root@localhost:62465/defaultdb?sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	for range time.NewTicker(time.Second).C {
		row := db.QueryRow(context.Background(), "SELECT version()")

		var version string
		if err = row.Scan(&version); err != nil {
			log.Printf("error scanning row: %v", err)
		}

		log.Println(version)
	}
}
