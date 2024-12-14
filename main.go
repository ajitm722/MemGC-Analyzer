package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

// Data type for memory allocation
type data struct {
	i, j int
}

// AllocateMemoryWithManualGC allocates memory and triggers garbage collection manually after each iteration
func AllocateMemoryWithManualGC(N int) {
	fmt.Println("Starting to allocate memory (with manual GC after each iteration)...")

	// Allocate memory by creating and discarding slices
	for i := 0; i < 3; i++ {
		structure := make([]data, 0, N) // Create a new slice
		for j := 0; j < N; j++ {        // Fill only a fraction of N for demo
			structure = append(structure, data{j, j})
		}
		fmt.Printf("Iteration %d: Allocated %d elements\n", i, len(structure))
		runtime.GC() // Trigger garbage collection manually after each iteration
		// The structure slice goes out of scope here, and the GC can clean it up
	}

	fmt.Println("Finished all iterations.")
	PrintMemoryStats()
}

// AllocateMemoryWithAutoGC allocates memory and triggers garbage collection manually at the end of all iterations
func AllocateMemoryWithAutoGC(N int) {
	fmt.Println("Starting to allocate memory (with auto GC at the end)...")

	// Allocate memory by creating and discarding slices
	for i := 0; i < 3; i++ {
		structure := make([]data, 0, N) // Create a new slice
		for j := 0; j < N; j++ {        // Fill only a fraction of N for demo
			structure = append(structure, data{j, j})
		}
		fmt.Printf("Iteration %d: Allocated %d elements\n", i, len(structure))
		// The structure slice goes out of scope here, and the GC can clean it up
	}

	fmt.Println("Finished all iterations.")
	PrintMemoryStats()
}

// PrintMemoryStats prints the current memory statistics
func PrintMemoryStats() {
	// Print memory stats before or after triggering garbage collection
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v MB, TotalAlloc = %v MB, HeapAlloc = %v MB\n",
		memStats.Alloc/1024/1024, memStats.TotalAlloc/1024/1024, memStats.HeapAlloc/1024/1024)
}

// Handler for HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Received request with N=%d and GC type=%s\n", N, gcType)

	// Determine which allocation function to call based on the GC type
	if strings.ToLower(gcType) == "manual" {
		fmt.Println("\nRunning AllocateMemoryWithManualGC...")
		AllocateMemoryWithManualGC(N)
	} else if strings.ToLower(gcType) == "auto" {
		fmt.Println("\nRunning AllocateMemoryWithAutoGC...")
		AllocateMemoryWithAutoGC(N)
	} else {
		http.Error(w, "Invalid 'gc' parameter, must be 'manual' or 'auto'", http.StatusBadRequest)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Memory allocation completed\n"))
}

// Main function to set up the server
func main() {
	// Register handler for HTTP requests
	http.HandleFunc("/allocate", handler)

	// Start the server
	fmt.Println("Server is starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
