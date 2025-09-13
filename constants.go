//go:build tinygo && (rp2040 || rp2350)

package tinygo_logger

var (
	// timestampBuffer is the buffer used for timestamp messages
	timestampBuffer = [8]byte{}

	// messageBuffer is the buffer used for log messages
	messageBuffer = [512]byte{}

	// messageIndex is the current index in the message buffer
	messageIndex = 0

	// debugHeader is the header for debug messages
	debugHeader = []byte("DEBUG")

	// warningHeader is the header for warning messages
	warningHeader = []byte("WARNING")

	// errorHeader is the header for error messages
	errorHeader = []byte("ERROR")

	// infoHeader is the default header for info messages
	infoHeader = []byte("INFO")

	// fullBufferHeader is the header for full buffer messages
	fullBufferHeader = []byte("FULL_BUFFER")

	// whitespaceBuffer is a byte slice representing a whitespace character
	whitespaceBuffer = []byte(" ")
	
	// newlineBuffer is a byte slice representing a newline character
	newlineBuffer = []byte("\n")

	// tabBuffer is a byte slice representing a tab character
	tabBuffer = []byte("\t")

	// twoPointsBuffer is a byte slice representing two points
	twoPointsBuffer = []byte(":")

	// dotBuffer is a byte slice representing a dot character
	dotBuffer = []byte(".")

	// hexPrefix is the prefix for error codes
	hexPrefix = []byte("0x")

	// float64Buffer is the buffer used for float64 messages
	float64Buffer = [8]byte{}

	// asciiHexDigits is a byte slice representing ASCII hex digits
	asciiHexDigits = []byte("0123456789ABCDEF")

	// asciiDecimalDigits is a byte slice representing ASCII decimal digits
	asciiDecimalDigits = []byte("0123456789")

	// uintToHexBuffer is a buffer used for converting uint64 to hex
	uintToHexBuffer = [16]byte{}

	// uintToDecimalBuffer is a buffer used for converting uint64 to decimal
	uintToDecimalBuffer = [20]byte{}
)