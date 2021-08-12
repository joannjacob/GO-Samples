package main

import (
	h "jwt_go_sample/handlers"
	"log"
	"net/http"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", h.Signin)
	http.HandleFunc("/welcome", h.Welcome)
	http.HandleFunc("/refresh", h.Refresh)

	// start the server on port 8001
	log.Fatal(http.ListenAndServe(":8001", nil))
}
