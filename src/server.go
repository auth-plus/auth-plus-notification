package main

import (
	http "auth-plus-notification/presentation/http"
	// kafka "auth-plus-notification/presentation/messaging"
)

func main() {
	// go kafka.Server()
	http.Server()
}
