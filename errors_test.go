package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCompatibility(t *testing.T) {
	t.Run("errors.Is compatibility", func(t *testing.T) {
		e1 := Wrap(nil, "foobar")
		e2 := Wrap(e1, "foobar 2")
		e3 := Wrap(e2, "foobar 3")
		e4 := Wrap(e2, "foobar 4")

		assert.True(t, errors.Is(e3, e1))
		assert.True(t, errors.Is(e3, e2))
		assert.True(t, errors.Is(e3, e3))
		assert.False(t, errors.Is(e3, e4))
		assert.True(t, errors.Is(e4, e1))

		someFunc := func() error {
			return Wrap(e1, "someFunc error")
		}

		someErr := someFunc()
		assert.True(t, errors.Is(someErr, e1))

		stringErr := errors.New("")
		var nilErr *Error
		var anotherNilErr *Error

		assert.True(t, errors.Is(nilErr, anotherNilErr))
		assert.False(t, errors.Is(nilErr, stringErr))
		assert.False(t, errors.Is(stringErr, nilErr))
	})

	t.Run("errors.Is compatibility", func(t *testing.T) {
		e1 := Wrap(nil, "foobar")
		e2 := Wrap(e1, "foobar 2")
		e3 := fmt.Errorf("foobar 3: %w", e2)

		var e4 *Error
		assert.True(t, errors.As(e3, &e4))
		assert.Equal(t, "foobar 2: foobar", e4.Error())

		customerErr1 := customErr{msg: "custom foobar"}
		e5 := Wrap(customerErr1, "foobar 5")
		e6 := Wrap(e5, "foobar 6")

		var customerErr2 customErr
		assert.True(t, errors.As(e6, &customerErr2))
		assert.Equal(t, "custom foobar", customerErr2.Error())
	})
}

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return ce.msg
}

func TestError_Contextual(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		e1 := Wrap(nil, "foobar",
			ContextualOption("foo", "bar"),
			ContextualOption("fooz", "barz"),
		)

		foo, fooExist := e1.Contextual("foo")
		fooz, foozExist := e1.Contextual("fooz")

		assert.True(t, fooExist)
		assert.Equal(t, "bar", foo)
		assert.True(t, foozExist)
		assert.Equal(t, "barz", fooz)
	})

	t.Run("key does not exist", func(t *testing.T) {
		e1 := Wrap(nil, "foobar",
			ContextualOption("foo", "bar"),
			ContextualOption("fooz", "barz"),
		)

		_, fooExist := e1.Contextual("fooo")

		assert.False(t, fooExist)
	})
}
