package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.SetFlags(0)

	db, err := pgxpool.New(context.Background(), "postgres://root@localhost:26257/defaultdb?sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	var i int
	for range time.NewTicker(time.Second).C {
		const stmt = `SELECT node_id, value
					  FROM crdb_internal.node_build_info
					  WHERE field = 'Version'`

		row := db.QueryRow(context.Background(), stmt)

		var nodeID int8
		var version string
		if err = row.Scan(&nodeID, &version); err != nil {
			log.Printf("error scanning row: %v", err)
		}

		log.Printf("node: %d\tversion: %s (iteration %d)", nodeID, version, i)
		i++
	}
}
