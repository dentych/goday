package hasher

import (
	"crypto/sha256"
	"fmt"
)

// Hash generates SHA256 hash for string s
func Hash(s string) string {
	fmt.Println("Hashing string:", s)
	hash := sha256.Sum256([]byte(s))
	niceHash := fmt.Sprintf("%x", hash)
	fmt.Println("Created hash:", hash)
	return niceHash
}
