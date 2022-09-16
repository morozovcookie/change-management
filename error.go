package cm

import (
	"bytes"
	"errors"
	"fmt"
)

var _ fmt.Stringer = (*ErrorCode)(nil)

// ErrorCode represents a type of the error.
type ErrorCode string

const (
	ErrorCodeOK       = ErrorCode("ok")
	ErrorCodeInvalid  = ErrorCode("invalid")
	ErrorCodeConflict = ErrorCode("conflict")
	ErrorCodeNotFound = ErrorCode("not_found")
	ErrorCodeInternal = ErrorCode("internal")
)

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (ec ErrorCode) String() string {
	return string(ec)
}

var _ error = (*Error)(nil)

// Error represents an error.
type Error struct {
	// Code is the machine-readable code, for reference purpose.
	Code ErrorCode

	// Message is the human-readable message for end user.
	Message string

	// Err is the embed error.
	Err error
}

func (err *Error) Error() string {
	if err == nil {
		return ""
	}

	if err.Err != nil {
		return err.Err.Error()
	}

	var buf bytes.Buffer

	_, _ = fmt.Fprintf(&buf, "<%s>", ErrorCodeFromError(err))

	if message := MessageFromError(err); message != "" {
		_, _ = fmt.Fprintf(&buf, " %s", message)
	}

	return buf.String()
}

func ErrorCodeFromError(err error) ErrorCode {
	if err == nil {
		return ErrorCodeOK
	}

	var customErr *Error
	ok := errors.As(err, &customErr)

	if ok && customErr.Code != "" {
		return customErr.Code
	}

	if ok && customErr.Err != nil {
		return ErrorCodeFromError(customErr.Err)
	}

	return ErrorCodeInternal
}

func MessageFromError(err error) string {
	if err == nil {
		return ""
	}

	var customErr *Error
	ok := errors.As(err, &customErr)

	if ok && customErr.Message != "" {
		return customErr.Message
	}

	if ok && customErr.Err != nil {
		return MessageFromError(customErr.Err)
	}

	return "an internal error has occurred"
}
