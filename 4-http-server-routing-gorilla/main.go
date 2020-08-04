package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi foo")
	}).Methods("GET").Host("localhost")

	router.HandleFunc("/users/{user}", func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods("GET")

	router.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
