package tinygo_logger

var (
	// DebugHeader is the header for debug messages
	DebugHeader = []byte("DEBUG")

	// WarningHeader is the header for warning messages
	WarningHeader = []byte("WARNING")

	// ErrorHeader is the header for error messages
	ErrorHeader = []byte("ERROR")

	// InfoHeader is the default header for info messages
	InfoHeader = []byte("INFO")

	// FullBufferHeader is the header for full buffer messages
	FullBufferHeader = []byte("FULL_BUFFER")
)
