package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dentych/goday/apprication/downloader"
	"github.com/gorilla/mux"
)

type postBody struct {
	Host     string
	Path     string
	Username string
	Password string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", post)
	http.ListenAndServe(":8080", router)
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body postBody
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Decode failed:", err)
	}

	downloader.DownloadFile(body.Host, body.Username, body.Username, body.Password)
}
