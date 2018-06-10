package james

import (
	"fmt"
)

// GoFile represents database row
type GoFile struct {
	hash     string
	location string
}

// Save persists strings hash and location to mssql database
func Save(hash, location string) bool {
	fmt.Println(hash, location)
	return false
}
