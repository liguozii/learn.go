package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/liguozi/go/src")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
