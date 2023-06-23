package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDataBase() {
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("POSTGRES_HOST", "localhost"),
		GetEnv("POSTGRES_PORT", "5432"),
		GetEnv("POSTGRES_USER", "khodza"),
		GetEnv("POSTGRES_PASSWORD", "1"),
		GetEnv("POSTGRES_DB", "rest-api-go"))

	var err error

	db, err = sqlx.Open("postgres", conString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Connected to the database")
}

func GetDB() *sqlx.DB {
	return db
}
