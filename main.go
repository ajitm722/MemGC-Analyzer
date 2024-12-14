package main

import (
	"fmt"
	"gc_check/handler" // Import the handler package
	"net/http"
)

func main() {
	// Register handler for HTTP requests
	http.HandleFunc("/allocate", handler.Handler)

	// Start the server
	fmt.Println("Server is starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
