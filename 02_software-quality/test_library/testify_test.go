package testify_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		person := &Person{FirstName: "sing"}
		if assert.NotNil(t, person) {
			assert.Equal(t, "sing", person.FirstName)
		}
	})

}
