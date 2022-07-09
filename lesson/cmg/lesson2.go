package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "hello get")
	} else {
		fmt.Fprintf(w, "other method")
	}
}
