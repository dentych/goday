package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dentych/goday/apprication/archiver"
	"github.com/dentych/goday/apprication/downloader"
	"github.com/gorilla/mux"
	"log"
)

type postBody struct {
	Host     string
	Path     string
	Username string
	Password string
}

// StartAPI starts the server
func StartAPI() {
	fmt.Println("Starting api...")
	router := mux.NewRouter()
	router.HandleFunc("/", post)
	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body postBody
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Decode failed:", err)
	}

	res := downloader.DownloadFile(body.Host, body.Username, body.Password, body.Path)
	defer res.Close()
	archiver.Archive(body.Path, res)
}
