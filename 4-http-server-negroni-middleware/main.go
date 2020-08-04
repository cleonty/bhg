package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type trivial struct {
}

func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing trivial middleware")
	next(w, r)
}
func main() {
	router := mux.NewRouter()
	middleware := negroni.Classic()

	middleware.UseHandler(router)
	middleware.Use(&trivial{})
	http.ListenAndServe(":8080", middleware)
}
