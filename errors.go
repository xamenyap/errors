package errors

import (
	"errors"
)

// Error wraps an inner error, allows adding contextual data to enrich the inner error
// and provides a friendly message that is safe for external system's consumption.
type Error struct {
	inner           error
	contextual      map[string]any
	FriendlyMessage string
}

// Error returns the full error message: the friendly message it contains
// and it's inner error message.
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

// Is asserts if error can be cast to target type.
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

	inner := e.inner
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

// Unwrap returns the inner error, or nil.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.inner
}

// Wrap wraps an inner error, along with a friendly message.
func Wrap(inner error, friendlyMessage string, options ...Option) *Error {
	e := &Error{
		inner:           inner,
		FriendlyMessage: friendlyMessage,
	}

	for _, opt := range options {
		opt(e)
	}

	return e
}

// Option is the optional parameter for Wrap.
type Option func(e *Error)

// ContextualOption is an Option that adds contextual data to Error using a key/value pair.
func ContextualOption(key string, value any) Option {
	return func(e *Error) {
		e.contextual[key] = value
	}
}
