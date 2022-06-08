package errors

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type _error struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

var (
	ErrorAuth       = &_error{Code: 0001, Message: "authentication error", Details: nil}
	ErrorPerm       = &_error{Code: 0002, Message: "permission error", Details: nil}
	ErrorValidation = &_error{Code: 0003, Message: "validation failed", Details: nil}
	ErrorNoResult   = &_error{Code: 0004, Message: "no result", Details: nil}
)

// Wrap 用來替覆寫已定義好的 error 中的 message
func Wrap(err error, message string) error {
	if tmp, ok := err.(*_error); ok {
		tmp.Message = message
		return tmp
	}
	return err
}

func (e *_error) Error() string {
	var b strings.Builder
	_, _ = b.WriteRune('[')
	_, _ = b.WriteString(strconv.Itoa(e.Code))
	_, _ = b.WriteRune(']')
	_, _ = b.WriteRune(' ')
	_, _ = b.WriteString(e.Message)
	return b.String()
}

func (e *_error) Is(target error) bool {
	causeErr := errors.Cause(target)
	tErr, ok := causeErr.(*_error)
	if !ok {
		return false
	}
	return e.Code == tErr.Code
}
