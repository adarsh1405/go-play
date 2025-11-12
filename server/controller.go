package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adarsh1405/go-play/connector"
)


func deleteAllEntry() {
	// Connect to Postgres DB
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	sqlStatement := `DELETE FROM employee;`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("failed to delete data: %v", err)
	}
	log.Println("All records deleted from the employee table successfully")
}

func insertAllEntry() {
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	for _, account := range accounts {
		sqlStatement := `
		INSERT INTO employee (id, name, username, email, company_name, company_catchphrase)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO NOTHING;`
		_, err = db.Exec(sqlStatement, account.ID, account.Name, account.Username, account.Email, account.Company.Name, account.Company.CatchPhrase)
		if err != nil {
			log.Fatalf("failed to insert data: %v", err)
		}
	}
	log.Println("All user details inserted into the database successfully")
}

func insertOneEntry(account account) {
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	sqlStatement := `
	INSERT INTO employee (id, name, username, email, company_name, company_catchphrase)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id) DO NOTHING;`
	_, err = db.Exec(sqlStatement, account.ID, account.Name, account.Username, account.Email, account.Company.Name, account.Company.CatchPhrase)
	if err != nil {
		log.Fatalf("failed to insert data: %v", err)
	}
	log.Println("User detail inserted into the database successfully")
}


func getAll() {
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	sqlStatement := `SELECT id, name, username, email, company_name, company_catchphrase FROM employee;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("failed to retrieve data: %v", err)
	}
	defer rows.Close()

	// clear the accounts slice before appending
	accounts = accounts[:0]

	for rows.Next() {
		var acc account
		err := rows.Scan(&acc.ID, &acc.Name, &acc.Username, &acc.Email, &acc.Company.Name, &acc.Company.CatchPhrase)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		accounts = append(accounts, acc)
	}

}

func getOne(id int) (account, error) {
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	var acc account
	sqlStatement := `SELECT id, name, username, email, company_name, company_catchphrase FROM employee WHERE id = $1;`
	err = db.QueryRow(sqlStatement, id).Scan(&acc.ID, &acc.Name, &acc.Username, &acc.Email, &acc.Company.Name, &acc.Company.CatchPhrase)
	if err != nil {
		if err == sql.ErrNoRows {
			return acc , fmt.Errorf("ID not found")
		}
		return acc, fmt.Errorf("failed to query row: %v", err)
	}
	return acc, nil
}