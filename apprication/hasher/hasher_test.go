package hasher

import (
	"fmt"
	"testing"
)

func TestArchive(t *testing.T) {
	test := "Hello"
	hash := Hash(test)

	fmt.Println(hash)
}
