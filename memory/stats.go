package memory

import (
	"fmt"
	"runtime"
)

// PrintMemoryStats prints the current memory statistics
func PrintMemoryStats() {
	// Print memory stats before or after triggering garbage collection
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v MB, TotalAlloc = %v MB, HeapAlloc = %v MB\n",
		memStats.Alloc/1024/1024, memStats.TotalAlloc/1024/1024, memStats.HeapAlloc/1024/1024)
}

