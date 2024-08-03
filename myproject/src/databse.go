package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/cetec")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
}

func getDB() *sql.DB {
	return db
}
