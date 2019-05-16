package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"template/types"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/example", GetExample).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	log.Print(err)
}

func GetExample(w http.ResponseWriter, r *http.Request) {
	example := types.Example{
		FirstName: "joseph",
		LastName:  "johnson",
	}
	jsonResponse(w, http.StatusOK, example)
}

func jsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		log.Print(err)
	}
}