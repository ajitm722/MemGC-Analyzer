package memory

import (
	"fmt"
	"runtime"
)

// AllocateMemoryWithManualGC allocates memory and triggers garbage collection manually after each iteration
func AllocateMemoryWithManualGC(N int) {
	fmt.Println("Starting to allocate memory (with manual GC after each iteration)...")

	// Allocate memory by creating and discarding slices
	for i := 0; i < 3; i++ {
		structure := make([]byte, 0, N) // Create a new slice of bytes
		for j := 0; j < N; j++ {
			structure = append(structure, byte(j%256)) // Add bytes N times (mod 256 to fit in a byte)
		}
		fmt.Printf("Iteration %d: Allocated %d bytes\n", i, len(structure))
		runtime.GC() // Trigger garbage collection manually after each iteration
		// The structure slice goes out of scope here, and the GC can clean it up
	}

	fmt.Println("Finished all iterations.")
}

// AllocateMemoryWithAutoGC allocates memory and triggers garbage collection at the end of all iterations
func AllocateMemoryWithAutoGC(N int) {
	fmt.Println("Starting to allocate memory (with auto GC after all iterations)...")

	// Allocate memory by creating and discarding slices
	for i := 0; i < 3; i++ {
		structure := make([]byte, 0, N) // Create a new slice of bytes
		for j := 0; j < N; j++ {
			structure = append(structure, byte(j%256)) // Add bytes N times (mod 256 to fit in a byte)
		}
		fmt.Printf("Iteration %d: Allocated %d bytes\n", i, len(structure))
		// The structure slice goes out of scope here, and the GC can clean it up
	}

	fmt.Println("Finished all iterations.")
}
