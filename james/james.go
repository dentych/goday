package james

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

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
