package main

import (
	"fmt"
	"net/http"

	"github.com/vidalpaul/bnb/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
