package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dentych/goday/apprication/downloader"
	"github.com/gorilla/mux"
	"github.com/dentych/goday/apprication/archiver"
)

type postBody struct {
	Host     string
	Path     string
	Username string
	Password string
}

func StartApi() {
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

	res := downloader.DownloadFile(body.Host, body.Username, body.Password, body.Path)
	defer res.Close()
	archiver.Archive(body.Path, res)
}
