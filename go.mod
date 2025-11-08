module github.com/adarsh1405/go-play

go 1.25.3

require (
	github.com/adarsh1405/go-play/connector v0.0.0
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

// Use local connector module path to resolve nested module
replace github.com/adarsh1405/go-play/connector => ./connector
