//go:build tinygo && (rp2040 || rp2350)

package tinygo_logger

import (
	tinygotypes "github.com/ralvarezdev/tinygo-types"
)

type (
	// Logger is an interface for logging messages
	Logger interface {
		AddSpace()
		AddNewline()
		AddTab()
		AddHexCode(hexCode []byte, newline bool)
		AddErrorCode(errCode tinygotypes.ErrorCode, newline bool)
		AddUint8(value uint8, newline bool, hexCode bool)
		AddUint16(value uint16, newline bool, hexCode bool)
		AddUint32(value uint32, newline bool, hexCode bool)
		AddUint64(value uint64, newline bool, hexCode bool)
		AddFloat64(value float64, newline bool)
		AddMessage(message []byte, newline bool)
		AddMessageWithHexCode(message []byte, hexBuffer []byte, separate bool, newline bool)
		AddMessageWithErrorCode(message []byte, errCode tinygotypes.ErrorCode, separate bool, newline bool)
		AddMessageWithUint8(message []byte, value uint8, separate bool, newline bool, hexCode bool)
		AddMessageWithUint16(message []byte, value uint16, separate bool, newline bool, hexCode bool)
		AddMessageWithUint32(message []byte, value uint32, separate bool, newline bool, hexCode bool)
		AddMessageWithUint64(message []byte, value uint64, separate bool, newline bool, hexCode bool)
		AddMessageWithFloat64(message []byte, value float64, separate bool, newline bool)
		Debug()
		DebugMessage(message []byte)
		Info()
		InfoMessage(message []byte)
		Warning()
		WarningMessage(message []byte)
		WarningMessageWithErrorCode(message []byte, errCode tinygotypes.ErrorCode, separate bool)
		Error()
		ErrorMessage(message []byte)
		ErrorMessageWithErrorCode(message []byte, errCode tinygotypes.ErrorCode, separate bool)
	}
)