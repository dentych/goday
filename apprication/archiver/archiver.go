package archiver

import (
	"github.com/jlaffaye/ftp"
	"io"
	"os"
	"log"
	"strings"
	"github.com/dentych/goday/apprication/hasher"
)

// Archive does stuff...
func Archive(filename string, response *ftp.Response) {
	splitFilename := strings.Split(filename, "/")
	filenameWithoutPath := splitFilename[len(splitFilename)-1]
	hash := hasher.Hash(filenameWithoutPath)
	file, err := os.Create(hash)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	log.Println("Writing file:", filenameWithoutPath, " as hash:", string(hash))
	io.Copy(file, response)

	// call james with hash and filenameWithoutPath
}
