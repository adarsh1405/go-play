package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func main() {
	data = GetUserData()
	decodedata(data)

	connector.Hello()


	db, err := connector.ConnectPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing db: %v", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/fetch", fetchDetails).Methods("GET")
	// r.HandleFunc("/insert", insertDetails).Methods("POST")
	// r.HandleFunc("/delete", deleteDetails).Methods("DELETE")

	
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func GetUserData() string {
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Response Type: %T\n", content). //content is of type []uint8 -
	// fmt.Println(content)

	return string(content)
}

func decodedata(data string) {

	valid := json.Valid([]byte(data))

	if valid {

		err := json.Unmarshal([]byte(data), &accounts)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("Decoded Data: %+v\n", accounts)
		// for _, account := range accounts {
		// 	fmt.Printf("ID: %d\nName: %s\nUsername: %s\nEmail: %s\nCompany Name: %s\nCatch Phrase: %s\n\n",
		// 		account.ID, account.Name, account.Username, account.Email, account.Company.Name, account.Company.CatchPhrase)
		// }
	} else {
		fmt.Println("The data is NOT valid JSON")
	}

}

func fetchDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
	log.Println("Status - " , http.StatusOK)
	
}
