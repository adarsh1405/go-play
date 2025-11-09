package connector

import (
	"database/sql"
	"fmt"
	"log"

	// register postgres driver
	_ "github.com/lib/pq"
)

const (
  	host     = "localhost"
  	port     = 5432
  	user     = "adarshpadhi"
  	password = ""
	dbname   = "employee_info"
)

// ConnectPostgresDB opens and verifies a connection to Postgres.
// It returns the *sql.DB for callers to use and close, or an error.
func ConnectPostgresDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// verify the connection is alive
	if err := db.Ping(); err != nil {
		// close opened handle on error to avoid leaks
		_ = db.Close()
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	log.Println("Connected to Postgres Database successfully!")
	fmt.Println("Connected to Postgres Database successfully!")
	return db, nil
}

func Hello() {
	fmt.Println("Hello from connector package")
}