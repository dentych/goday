package james

import (
	"testing"
)

func TestSave(t *testing.T) {
	hash := "ASD123"
	location := "some/path/to/stuff"

	result := Save(hash, location)

	if result == false {
		t.Errorf("You done goofed, son!")
	}
}
