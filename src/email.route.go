package main

import (
	"fmt"
	"net/http"
)

func EmailHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Ok\n")
}
