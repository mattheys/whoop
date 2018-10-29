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
	b := Response{Status: "YesIAmNew"}
	json.NewEncoder(w).Encode(b)
}

// GetNew2 returns always the same response.
func GetNew2(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "YesIAmNew2"}
	json.NewEncoder(w).Encode(b)
}

// GetNew3 returns always the same response.
func GetNew3(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "YesIAmNew3"}
	json.NewEncoder(w).Encode(b)
}

// GetNew4 returns always the same response.
func GetNew4(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "YesIAmNew3"}
	json.NewEncoder(w).Encode(b)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", GetStatus).Methods("GET")
	router.HandleFunc("/new", GetNew).Methods("GET")
	router.HandleFunc("/new2", GetNew2).Methods("GET")
	router.HandleFunc("/new3", GetNew3).Methods("GET")
	router.HandleFunc("/new4", GetNew4).Methods("GET")
	log.Fatal(http.ListenAndServe(":8123", router))
}
