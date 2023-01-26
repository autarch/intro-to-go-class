package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert" // HL
)

func TestNew(t *testing.T) {
	user, err := New("", "password")

	assert.Error(t, err, "got error from constructor") // HL
	assert.EqualError(                                       // HL
		t,
		err,
		"The username must be a non-empty string.",
		"got the expected error for an empty username",
	)
	assert.Equal(t, User{}, user, "user returned on error is empty") // HL
}
