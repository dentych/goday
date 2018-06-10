package archiver

import (
	"bytes"
	"encoding/json"
	"github.com/dentych/goday/apprication/hasher"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type postBody struct {
	Hash     string
	Filename string
}

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
	body := postBody{
		Hash:     hash,
		Filename: filenameWithoutPath,
	}
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	url := os.Getenv("JAMES_URL")
	if len(url) == 0 {
		url = "http://localhost:8081"
	}
	_, err = http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
}
