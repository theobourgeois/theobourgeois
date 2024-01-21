package db

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


var db * sql.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file", err)
	}

	// Initialize a new connection object to database
	cfg := mysql.Config{
		User:  	os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"), 
		Net:    "tcp",
		Addr:   "127.0.0.1:" + os.Getenv("DB_PORT"), 
		DBName: os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Database Connected Successfully on port 8889...")
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return db.Query(query)
}

func GetDB() *sql.DB {
	return db
}