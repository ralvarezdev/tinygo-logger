package tinygo_logger

import (
	"os"
	"time"

	tinygobuffers "github.com/ralvarezdev/tinygo-buffers"
	tinygoerrors "github.com/ralvarezdev/tinygo-errors"
)

type (
	// DefaultLogger is a simple implementation of the Logger interface
	DefaultLogger struct {
		messageBuffer []byte
		messageIndex  int
	}
)

// NewDefaultLogger creates a new DefaultLogger instance
//
// Parameters:
//
//	bufferSize: The size of the message buffer to allocate.
//
// Returns:
//
//	A pointer to the newly created DefaultLogger instance.
func NewDefaultLogger(bufferSize uint64) *DefaultLogger {
	return &DefaultLogger{
		messageBuffer: make([]byte, bufferSize),
	}
}

// writeTimestamp is a helper function to print the current timestamp
func (l *DefaultLogger) writeTimestamp() {
	now := time.Now().UnixNano() / int64(time.Millisecond)

	// Get the hour, minute, second, and millisecond components
	hour := now / time.Hour.Milliseconds()
	minute := (now % time.Hour.Milliseconds()) / time.Minute.Milliseconds()
	second := (now % time.Minute.Milliseconds()) / time.Second.Milliseconds()
	millisecond := now % time.Second.Milliseconds()

	// Print the timestamp in the format HH:MM:SS.mmm
	buffer := tinygobuffers.UintToDecimalFixed(uint64(hour), 2)
	os.Stdout.Write(buffer)
	os.Stdout.Write(tinygobuffers.TwoPointsBuffer)
	buffer = tinygobuffers.UintToDecimalFixed(uint64(minute), 2)
	os.Stdout.Write(buffer)
	os.Stdout.Write(tinygobuffers.TwoPointsBuffer)
	buffer = tinygobuffers.UintToDecimalFixed(uint64(second), 2)
	os.Stdout.Write(buffer)
	os.Stdout.Write(tinygobuffers.DotBuffer)
	buffer = tinygobuffers.UintToDecimalFixed(uint64(millisecond), 3)
	os.Stdout.Write(buffer)
}

// writeNewline is a helper function to print a newline
func (l *DefaultLogger) writeNewline() {
	os.Stdout.Write(tinygobuffers.NewlineBuffer)
}

// writeSpace is a helper function to print a space
func (l *DefaultLogger) writeSpace() {
	os.Stdout.Write(tinygobuffers.WhitespaceBuffer)
}

// writeHeader is a helper function to print the header if required
//
// Parameters:
//
//	header: Whether to include the header in the log message.
func (l *DefaultLogger) writeHeader(header []byte) {
	if header != nil {
		l.writeTimestamp()
		l.writeSpace()
		os.Stdout.Write(header)
		l.writeSpace()
	}
}

// writeMessage is a helper function to print the message from the messageBuffer
func (l *DefaultLogger) writeMessage() {
	if l.messageIndex > 0 {
		os.Stdout.Write(l.messageBuffer[:l.messageIndex])
		l.messageIndex = 0 // Reset index after printing
	}
}

// checkIndex checks if the messageIndex exceeds the messageBuffer size
func (l *DefaultLogger) checkIndex() {
	if l.messageIndex >= len(l.messageBuffer) {
		// Buffer full, print the message and reset the index
		l.log(FullBufferHeader)
		l.messageIndex = 0 // Reset index if it exceeds buffer size
	}
}

// AddSpace function to add a whitespace character to the messageBuffer
func (l *DefaultLogger) AddSpace() {
	l.messageBuffer[l.messageIndex] = tinygobuffers.WhitespaceBuffer[0]
	l.messageIndex++
	l.checkIndex()
}

// AddNewline function to add a newline character to the messageBuffer
func (l *DefaultLogger) AddNewline() {
	l.messageBuffer[l.messageIndex] = tinygobuffers.NewlineBuffer[0]
	l.messageIndex++
	l.checkIndex()
}

// AddTab function to add a tab character to the messageBuffer
func (l *DefaultLogger) AddTab() {
	l.messageBuffer[l.messageIndex] = tinygobuffers.TabBuffer[0]
	l.messageIndex++
	l.checkIndex()
}

// AddHexCode function to add hex code to the messageBuffer
//
// Parameters:
//
//	hexBuffer: The byte slice representing the hex code to print in hexadecimal format.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddHexCode(hexBuffer []byte, newline bool) {
	if hexBuffer != nil {
		for c := range tinygobuffers.HexPrefix {
			l.messageBuffer[l.messageIndex] = tinygobuffers.HexPrefix[c]
			l.messageIndex++
			l.checkIndex()
		}
		for c := range hexBuffer {
			l.messageBuffer[l.messageIndex] = hexBuffer[c]
			l.messageIndex++
			l.checkIndex()
		}

		if newline {
			l.AddNewline()
		}
	}
}

// AddErrorCode function to add an error code to the messageBuffer
//
// Parameters:
//
//	errCode: The error code to add to the message buffer.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddErrorCode(
	errCode tinygoerrors.ErrorCode,
	newline bool,
) {
	l.AddUint16(uint16(errCode), newline, true)
}

// AddUint8 function to add an uint8 value to the messageBuffer
//
// Parameters:
//
//	value: The uint8 value to add.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint8 value in hexadecimal format.
func (l *DefaultLogger) AddUint8(value uint8, newline bool, hexCode bool) {
	if hexCode {
		buffer := tinygobuffers.Uint8ToHex(value)
		l.AddHexCode(buffer, newline)
	} else {
		buffer := tinygobuffers.UintToDecimal(uint64(value))
		l.AddMessage(buffer, newline)
	}
}

// AddUint16 function to add an uint16 value to the messageBuffer
//
// Parameters:
//
//	value: The uint16 value to add.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint16 value in hexadecimal format.
func (l *DefaultLogger) AddUint16(value uint16, newline bool, hexCode bool) {
	if hexCode {
		buffer := tinygobuffers.Uint16ToHex(value)
		l.AddHexCode(buffer, newline)
	} else {
		buffer := tinygobuffers.UintToDecimal(uint64(value))
		l.AddMessage(buffer, newline)
	}
}

// AddUint32 function to add an uint32 value to the messageBuffer
//
// Parameters:
//
//	value: The uint32 value to add.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint32 value in hexadecimal format.
func (l *DefaultLogger) AddUint32(value uint32, newline bool, hexCode bool) {
	if hexCode {
		buffer := tinygobuffers.Uint32ToHex(value)
		l.AddHexCode(buffer, newline)
	} else {
		buffer := tinygobuffers.UintToDecimal(uint64(value))
		l.AddMessage(buffer, newline)
	}
}

// AddUint64 function to add an uint64 value to the messageBuffer
//
// Parameters:
//
//	value: The uint64 value to add.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint64 value in hexadecimal format.
func (l *DefaultLogger) AddUint64(value uint64, newline bool, hexCode bool) {
	if hexCode {
		buffer := tinygobuffers.Uint64ToHex(value)
		l.AddHexCode(buffer, newline)
	} else {
		buffer := tinygobuffers.UintToDecimal(uint64(value))
		l.AddMessage(buffer, newline)
	}
}

// AddFloat64 function to add a float64 value to the messageBuffer
//
// Parameters:
//
//	value: The float64 value to add.
//	precision: The number of decimal places to include.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddFloat64(value float64, precision int, newline bool) {
	for i := precision; i >= 0; i-- {
		// Store the float64 value in the buffer
		buffer, err := tinygobuffers.Float64ToDecimal(value, i)
		if err != tinygoerrors.ErrorCodeNil {
			continue
		}
		l.AddMessage(buffer, newline)
		return
	}

	// If all precisions fail, fallback to raw float64 representation
	tinygobuffers.Float64ToBytes(value, tinygobuffers.Float64Buffer[:])
	l.AddMessage(tinygobuffers.Float64Buffer[:], newline)
}

// AddMessage function to add a message to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddMessage(message []byte, newline bool) {
	if message != nil {
		for c := range message {
			l.messageBuffer[l.messageIndex] = message[c]
			l.messageIndex++
			l.checkIndex()
		}

		if newline {
			l.AddNewline()
		}
	}
}

// AddMessageWithHexCode function to add a message and hex code to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	hexBuffer: The byte slice representing the hex code to add in hexadecimal format.
//	separate: Whether to include a space between the message and hex code.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddMessageWithHexCode(
	message []byte,
	hexBuffer []byte,
	separate bool,
	newline bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddHexCode(hexBuffer, newline)
}

// AddMessageWithErrorCode function to add a message and error code to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	errCode: The error code to add to the message buffer.
//	separate: Whether to include a space between the message and error code.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddMessageWithErrorCode(
	message []byte,
	errCode tinygoerrors.ErrorCode,
	separate bool,
	newline bool,
) {
	l.AddMessageWithUint16(message, uint16(errCode), separate, newline, true)
}

// AddMessageWithUint8 function to add a message and uint8 value to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	value: The uint8 value to add.
//	separate: Whether to include a space between the message and uint8 value.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint8 value in hexadecimal format.
func (l *DefaultLogger) AddMessageWithUint8(
	message []byte,
	value uint8,
	separate bool,
	newline bool,
	hexCode bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddUint8(value, newline, hexCode)
}

// AddMessageWithUint16 function to add a message and uint16 value to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	value: The uint16 value to add.
//	separate: Whether to include a space between the message and uint16 value.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint16 value in hexadecimal format.
func (l *DefaultLogger) AddMessageWithUint16(
	message []byte,
	value uint16,
	separate bool,
	newline bool,
	hexCode bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddUint16(value, newline, hexCode)
}

// AddMessageWithUint32 function to add a message and uint32 value to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	value: The uint32 value to add.
//	separate: Whether to include a space between the message and uint32 value.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint32 value in hexadecimal format.
func (l *DefaultLogger) AddMessageWithUint32(
	message []byte,
	value uint32,
	separate bool,
	newline bool,
	hexCode bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddUint32(value, newline, hexCode)
}

// AddMessageWithUint64 function to add a message and uint64 value to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	value: The uint64 value to add.
//	separate: Whether to include a space between the message and uint64 value.
//	newline: Whether to include a newline at the end of the log message.
//	hexCode: Whether to add the uint64 value in hexadecimal format.
func (l *DefaultLogger) AddMessageWithUint64(
	message []byte,
	value uint64,
	separate bool,
	newline bool,
	hexCode bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddUint64(value, newline, hexCode)
}

// AddMessageWithFloat64 function to add a message and float64 value to the messageBuffer
//
// Parameters:
//
//	message: The byte slice representing the message to add.
//	value: The float64 value to add.
//	precision: The number of decimal places to include.
//	separate: Whether to include a space between the message and float64 value.
//	newline: Whether to include a newline at the end of the log message.
func (l *DefaultLogger) AddMessageWithFloat64(
	message []byte,
	value float64,
	precision int,
	separate bool,
	newline bool,
) {
	l.AddMessage(message, false)
	if separate {
		l.AddSpace()
	}
	l.AddFloat64(value, precision, newline)
}

// log functions for different log levels
//
// Parameters:
//
// header: The byte slice representing the log header to use.
func (l *DefaultLogger) log(header []byte) {
	l.writeHeader(header)
	l.writeMessage()
}

// Debug function to print debug messages with messageBuffer content
func (l *DefaultLogger) Debug() {
	l.log(DebugHeader)
}

// Warning function to print warning messages with messageBuffer content
func (l *DefaultLogger) Warning() {
	l.log(WarningHeader)
}

// Error function to print error messages with messageBuffer content
func (l *DefaultLogger) Error() {
	l.log(ErrorHeader)
}

// Info function to print info messages with messageBuffer content
func (l *DefaultLogger) Info() {
	l.log(InfoHeader)
}

// DebugMessage function to print a debug message
//
// Parameters:
//
//	message: The byte slice representing the debug message to print.
func (l *DefaultLogger) DebugMessage(message []byte) {
	l.AddMessage(message, true)
	l.Debug()
}

// WarningMessage function to print a warning message
//
// Parameters:
//
//	message: The byte slice representing the warning message to print.
func (l *DefaultLogger) WarningMessage(message []byte) {
	l.AddMessage(message, true)
	l.Warning()
}

// ErrorMessage function to print an error message
//
// Parameters:
//
//	message: The byte slice representing the error message to print.
func (l *DefaultLogger) ErrorMessage(message []byte) {
	l.AddMessage(message, true)
	l.Error()
}

// InfoMessage function to print an info message
//
// Parameters:
//
//	message: The byte slice representing the info message to print.
func (l *DefaultLogger) InfoMessage(message []byte) {
	l.AddMessage(message, true)
	l.Info()
}

// WarningMessageWithErrorCode function to print a warning message with an error code
//
// Parameters:
//
//	message: The byte slice representing the warning message to print.
//	errCode: The error code to add to the message buffer.
//	separate: Whether to include a space between the message and error code.
func (l *DefaultLogger) WarningMessageWithErrorCode(
	message []byte,
	errCode tinygoerrors.ErrorCode,
	separate bool,
) {
	l.AddMessageWithErrorCode(message, errCode, separate, true)
	l.Warning()
}

// ErrorMessageWithErrorCode function to print an error message with an error code
//
// Parameters:
//
//	message: The byte slice representing the error message to print.
//	errCode: The error code to add to the message buffer.
//	separate: Whether to include a space between the message and error code.
func (l *DefaultLogger) ErrorMessageWithErrorCode(
	message []byte,
	errCode tinygoerrors.ErrorCode,
	separate bool,
) {
	l.AddMessageWithErrorCode(message, errCode, separate, true)
	l.Error()
}
