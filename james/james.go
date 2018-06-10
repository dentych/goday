package james

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Save persists strings hash and location to mssql database
func Save(hash, location string) bool {

	db, err := sql.Open("mssql", "server=(localdb)\\mssqllocaldb;database=go")
	if err != nil {
		fmt.Println("Unable to connect to db:", err)
		return false
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Pinging db failed:", err)
		return false
	}
	defer db.Close()

	return false
}
