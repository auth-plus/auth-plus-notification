package main

import (
	http "auth-plus-notification/api/http"
	// kafka "auth-plus-notification/api/messaging"
)

func main() {
	// go kafka.Server()
	http.Server()
}
