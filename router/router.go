package router

import (
	"url_shortener/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", handlers.ShortenURLHandler).Methods("POST")
	router.HandleFunc("/{shortURL}", handlers.RedirectHandler).Methods("GET")
	return router
}
