package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"template/types"
)

type ExampleLister struct {
	examples []types.Example
}

func main() {
	exampler := ExampleLister{}
	r := mux.NewRouter()

	r.HandleFunc("/example", exampler.PostExample).Methods("POST")
	r.HandleFunc("/examples", exampler.GetExample).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	log.Print(err)
}

func (y *ExampleLister) PostExample(w http.ResponseWriter, r *http.Request) {
	var example types.Example
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&example)
	if err != nil {
		log.Print("submitted example could not be decoded")
		return
	}

	y.examples = append(y.examples, example)

	jsonResponse(w, http.StatusOK, y.examples)
}

func(y *ExampleLister) GetExample(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, y.examples)
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