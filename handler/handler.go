package handler

import (
	"fmt"
	"gc_check/memory" // Import the memory package
	"net/http"
	"strconv"
	"strings"
)

// Handler for HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters for N and GC type
	queryParams := r.URL.Query()
	nStr := queryParams.Get("N")
	gcType := queryParams.Get("gc")

	// Parse N from query parameter
	N, err := strconv.Atoi(nStr)
	if err != nil || N <= 0 {
		http.Error(w, "Invalid or missing 'N' parameter", http.StatusBadRequest)
		return
	}

	// Print the received parameters for logging purposes
	fmt.Printf("\n\nReceived request with N=%d and GC type=%s\n", N, gcType)

	// Determine which allocation function to call based on the GC type
	if strings.ToLower(gcType) == "manual" {
		fmt.Println("Running AllocateMemoryWithManualGC...")
		memory.AllocateMemoryWithManualGC(N)
		memory.PrintMemoryStats() // Print memory stats immediately after manual GC allocation
	} else if strings.ToLower(gcType) == "auto" {
		fmt.Println("Running AllocateMemoryWithAutoGC...")
		memory.AllocateMemoryWithAutoGC(N)
		memory.PrintMemoryStats() // Print memory stats immediately after auto GC allocation
	} else {
		http.Error(w, "Invalid 'gc' parameter, must be 'manual' or 'auto'", http.StatusBadRequest)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Memory allocation completed\n"))
}
