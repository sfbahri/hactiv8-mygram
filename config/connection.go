package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "127.0.0.1"
	port   = 5432
	user   = "postgres"
	dbname = "db_mygram"
)

func ConnDB() (*sql.DB, error) {
	// connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v", err)
	// 	return nil, err
	// }

	// return db, nil

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// Optional: Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
		return nil, err
	}

	return db, nil

}
