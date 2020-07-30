package main

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintln(w, "Executing /a")
	case "/b":
		fmt.Fprintln(w, "Executing /b")
	case "/c":
		fmt.Fprintln(w, "Executing /c")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8080", &r)
}
