package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Response is just a very basic example.
type Response struct {
	Status string `json:"status,omitempty"`
}

// GetStatus returns always the same response.
func GetStatus(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "idle"}
	json.NewEncoder(w).Encode(b)
}

// GetNew returns always the same response.
func GetNew(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "5"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", GetStatus).Methods("GET")
	router.HandleFunc("/new", GetNew).Methods("GET")
	log.Fatal(http.ListenAndServe(":8123", router))
}
