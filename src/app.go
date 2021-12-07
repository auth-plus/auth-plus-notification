package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func appHandler(res http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()
	s := r.PathPrefix("/").Subrouter()
	s.HandleFunc("/email", EmailHandler)
}
