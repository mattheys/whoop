package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

// GetStatus returns always the same response.
func GetNew(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "YesIAmNew"}
	json.NewEncoder(w).Encode(b)
}

func GetNew2(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "YesIAmNew2"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", GetStatus).Methods("GET")
	router.HandleFunc("/new", GetNew).Methods("GET")
	router.HandleFunc("/new2", GetNew2).Methods("GET")
	log.Fatal(http.ListenAndServe(":8123", router))
}
