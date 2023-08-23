package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {

	err := godotenv.Load()
	HandleFatal(err)

	GetEnv := func(key string) string {
		return os.Getenv(key)
	}

	connStr := "user=" + GetEnv("DB_USER") + " password=" + GetEnv("DB_PASSWORD") + " dbname=" + GetEnv("DB_NAME") + " sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	HandleFatal(err)

	err = db.Ping()
	HandleFatal(err)

}

func HandleFatal(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
