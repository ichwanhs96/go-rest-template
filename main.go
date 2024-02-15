package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func main() {
	http.HandleFunc("/hello", handleRequest)
	http.ListenAndServe(":8090", nil)
}
