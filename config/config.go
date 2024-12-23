package config

import (
	// "fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	connStr := os.Getenv("DB_CONNECTION_STRING")
	if connStr == "" {
		connStr = "postgres://postgres:Arosel420$@localhost/mydb?sslmode=disable"
	}
	// DB, err = sqlx.Connect("postgres", "user=postgres password=Arosel420$ dbname=mydb sslmode=disable")
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Ping error: %v", err)
	}
	// fmt.Println("Database connected successfully")
	log.Println("Starting server on :8080")
}
