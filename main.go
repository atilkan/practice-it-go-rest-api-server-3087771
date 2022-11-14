package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":9003", nil))
}
