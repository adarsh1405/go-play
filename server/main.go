package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/insert", insertDetails).Methods("POST")
	r.HandleFunc("/entries", getAllEntries).Methods("GET")
	r.HandleFunc("/id/{id}", getByID).Methods("GET")
	// r.HandleFunc("/delete", deleteDetails).Methods("DELETE")


		// Available endpoints
	log.Println("------------------------------")
	log.Println("Available endpoints:")
	log.Println("GET     /fetch      - fetch all user details")
	log.Println("POST    /insert     - insert user details")
	log.Println("GET     /entries    - get all user details from database")
	log.Println("GET     /id/{id}   - get user detail by ID from database")
	log.Println("DELETE  /delete     - delete user details")
	log.Println("------------------------------")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
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

func insertDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the single user detail from request body
	if err := json.NewDecoder(r.Body).Decode(&singleData); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	} 
	
	insertOneEntry(singleData)

	if err := json.NewEncoder(w).Encode("User detail inserted into the database successfully"); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
	log.Println("Status - " , http.StatusOK , " - Inserted user detail successfully")
}

func getAllEntries( w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	getAll()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
	log.Println("Status - " , http.StatusOK , " - Fetched all user details from database successfully")
}	

func getByID( w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res,err := getOne(intID)
	if err != nil {
		http.Error(w, "failed to get user detail by ID", http.StatusInternalServerError)
		log.Println("Error - " , err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
	log.Println("Status - " , http.StatusOK , " - Fetched user detail by ID from database successfully")
}	