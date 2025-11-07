package connector

import (
	"database/sql"
	"fmt"
	"log"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "your-password"
  dbname   = "calhounio_demo"
)

func connectPostgresDB() {
	// code to connect to Postgres database

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
	log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} 


}