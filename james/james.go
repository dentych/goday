package james

import (
	"database/sql"
	"fmt"
)

// GoFile represents database row
type GoFile struct {
	hash     string
	location string
}

// Save persists strings hash and location to mssql database
func Save(hash, location string) bool {
	db, err := sql.Open("mssql", "localhost")
	if err != nil {
		fmt.Println("Unable to connect to db", err)
	}
	rows := db.QueryRow("SELECT * FROM dbo.GoFiles")

	file := new(GoFile)
	rows.Scan(&file.hash, &file.location)

	fmt.Println(hash, location)
	return false
}
