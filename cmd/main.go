package main

import (
	"log"
	"net/http"
	"url_shortener/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
