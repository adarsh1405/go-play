package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adarsh1405/go-play/connector"
	"github.com/gorilla/mux"
)

const URL = "https://jsonplaceholder.typicode.com/users"

type account struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Company  struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
	} `json:"company"`

}


var accounts []account
var singleData account
var data string


func CheckConfigs() {

	//Connect to Postgres DB
	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()
	
}


func Run() {
	// Setting up the Router
	r := mux.NewRouter()
	r.HandleFunc("/fetch", fetchDetails).Methods("GET")
	// r.HandleFunc("/insert", insertDetails).Methods("POST")
	// r.HandleFunc("/delete", deleteDetails).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	// Available endpoints
	log.Println("Available endpoints:")
	log.Println("GET     /fetch      - fetch all user details")
	log.Println("POST    /insert     - insert user details")
	log.Println("DELETE  /delete     - delete user details")
}


func fetchDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// clear all values of accounts slice
	accounts = accounts[:0]

	data = GetUserData()
	Decodedata(data)

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
	log.Println("Status - " , http.StatusOK , " - Fetched all user details successfully")

	deleteAllEntry()
	insertAllEntry()
	
}


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
