package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/example/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		_, err := fmt.Fprintf(w, "You've requested the name %s", name)
		log.Print(err)
	})

	err := http.ListenAndServe(":80", r)
	log.Print(err)
}