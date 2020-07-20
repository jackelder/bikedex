package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	// Local DB
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		port, host, user, password, dbname)

	fmt.Println(connectionString)

	return sql.Open("postgres", connectionString)
}

func PingDatabase(db *sql.DB) error {
	return db.Ping()
}
