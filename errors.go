package errors

import (
	"errors"
)

type Error struct {
	Inner           error
	FriendlyMessage string
	ContextualData  map[string]interface{}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	msg := e.FriendlyMessage

	unwrapped := errors.Unwrap(e)
	for unwrapped != nil {
		msg += ": " + unwrapped.Error()
		unwrapped = errors.Unwrap(unwrapped)
	}

	return msg
}

func (e *Error) Is(target error) bool {
	tErr, ok := target.(*Error)
	if !ok {
		return false
	}

	if e == nil {
		return target == nil
	}

	if e.Error() == tErr.Error() {
		return true
	}

	inner := e.Inner
	for {
		if inner == nil {
			return false
		}

		if inner.Error() == tErr.Error() {
			return true
		}

		inner = errors.Unwrap(inner)
	}
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Inner
}

func Wrap(inner error, friendlyMessage string, contextualData map[string]interface{}) *Error {
	return &Error{
		Inner:           inner,
		FriendlyMessage: friendlyMessage,
		ContextualData:  contextualData,
	}
}
