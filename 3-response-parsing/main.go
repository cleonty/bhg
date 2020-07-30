package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const url = "https://www.google.com/robots.txt"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to get %s: %v", url, err)
	}
	fmt.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("unable to read response body for %s: %v", url, err)
	}
	fmt.Println(string(body))
}
