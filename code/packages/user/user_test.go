package user

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u1 := NewUser("bob")
	assert.Equal(t, u1.username, "bob", "u1 username is bob")

	u2 := NewUser("bob")
	assert.Equal(t, u2.username, "bob", "u2 username is bob")

	assert.Equal(
		t,
		reflect.ValueOf(u1).Pointer(),
		reflect.ValueOf(u2).Pointer(),
		"u1 and u2 point to the same struct",
	)
}
