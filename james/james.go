package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"
	"github.com/gorilla/mux"
	"log"

	_ "github.com/lib/pq"
)

type PostBody struct {
	Hash     string
	Filename string
}

// StartAPI starts the James API
func main() {
	fmt.Println("Starting api...")
	router := mux.NewRouter()
	router.HandleFunc("/", post)
	fmt.Println("Server listening on port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body PostBody
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Decode failed:", err)
	}

	result := Save(body.Hash, body.Filename)
	if result == false {
		fmt.Println("Failed to persist element:", body)
	}
}

// Save persists strings hash and location to mssql database
func Save(hash, filename string) bool {

	db, err := sql.Open("postgres", "host=tychsen.me user=postgres password=secretpassword dbname=goday sslmode=disable")
	if err != nil {
		log.Fatal("Couldn't open DB connection:", err)
		return false
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
		return false
	}
	defer db.Close()

	query := `INSERT INTO files (hash, filename) VALUES ($1, $2)`
	_, err = db.Exec(query, hash, filename)
	if err != nil {
		log.Fatal("Could not insert into database:", err)
		return false
	}

	return true
}
