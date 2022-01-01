package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
