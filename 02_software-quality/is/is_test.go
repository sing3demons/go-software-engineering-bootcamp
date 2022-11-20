package is_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func Binary(b string) (bool, error) {
	if b == "" {
		return false, errors.New("empty")
	}
	return true, nil
}

func TestSomething(t *testing.T) {
	is := is.New(t)

	b, err := Binary("1")

	is.NoErr(err)
	is.Equal(b, true)
	is.Equal([]string{"a", "b"}, []string{"a", "b"})

	got := "sing is gopher"
	is.True(strings.Contains(got, "sing"))

}
