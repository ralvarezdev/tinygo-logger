//go:build tinygo && (rp2040 || rp2350)

package tinygo_logger

import (
	"runtime"
)

var (
	// memoryStatsHeader is the header for memory statistics output
	memoryStatsHeader = []byte("Memory Stats:")

	// memoryStatsAllocKey is the key for currently allocated memory in KB
	memoryStatsAllocKey = []byte("\tCurrently Allocated (KB) =")

	// memoryStatsTotalAllocKey is the key for cumulative allocated memory in KB
	memoryStatsTotalAllocKey = []byte("\tTotal Allocated (KB) =")
)

// DebugMemory logs the current memory statistics using the provided Logger instance.
//
// Parameters:
//
// logger: The Logger instance to use for logging memory statistics.
func DebugMemory(logger Logger) {
	// Check if logger is nil
	if logger == nil {
		return
	}

	// Read memory statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

	// Log memory stats
	logger.AddMessage(memoryStatsHeader, true)
	logger.AddMessageWithUint64(memoryStatsAllocKey, m.Alloc/1024, true, true, false)
	logger.AddMessageWithUint64(memoryStatsTotalAllocKey, m.TotalAlloc/1024, true, true, false)
	logger.Debug()
}