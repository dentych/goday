package james

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type postBody struct {
	Hash     string
	Filename string
}

// StartAPI starts the James API
func StartAPI() {
	fmt.Println("Starting api...")
	router := mux.NewRouter()
	router.HandleFunc("/", post)
	fmt.Println("Server listening on port :8081")
	http.ListenAndServe(":8081", router)
}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body postBody
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Decode failed:", err)
	}

	result := Save(body.Hash, body.Filename)
	if result == false {
		fmt.Println("Failed to persist element:", body)
	}
}
