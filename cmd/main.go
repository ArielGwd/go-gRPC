package main

import (
	"flag"
	"log"
	"os"
	"proyek/pkg/database"
	"proyek/schema"

	_ "github.com/lib/pq"
)

func main() {
	// =========================================================================
	// App Starting

	log.Printf("main : Started")
	defer log.Println("main : Completed")

	// =========================================================================

	// Start Database
	db, err := database.OpenDatabase()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	// Handle cli command
	flag.Parse()

	if flag.Arg(0) == "migrate" {
		if err := schema.Migrate(db); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return
	}
}
