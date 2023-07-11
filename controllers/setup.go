package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/notes", GetAllNotes).Methods("GET")
	router.HandleFunc("/note/{id}", GetNote).Methods("GET")
	router.HandleFunc("/note", CreateNote).Methods("POST")
	router.HandleFunc("/note/{id}", UpdateNote).Methods("PUT")
	router.HandleFunc("/note/{id}", DeleteNote).Methods("DELETE")

	return router
}
