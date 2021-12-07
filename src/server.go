package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Ok\n")
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/health", healthHandler)
	route.HandleFunc("/", appHandler)
	http.Handle("/", route)
	log.Println("Listening on port 5000")
	http.ListenAndServe(":5000", route)
}
