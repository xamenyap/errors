package errors

import (
	"errors"
)

type Error struct {
	inner           error
	contextual      map[string]any
	FriendlyMessage string
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

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.inner
}

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

type Option func(e *Error)

func ContextualOption(key string, value any) Option {
	return func(e *Error) {
		e.contextual[key] = value
	}
}
