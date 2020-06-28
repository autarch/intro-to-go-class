package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert" // HL
)

func TestNew(t *testing.T) {
	user, err := NewUser("", "password")

	assert.NotNil(t, err, "got an error when the username is empty") // HL
	assert.Equal(                                                    // HL
		t,
		errors.New("The username must be a non-empty string."),
		err,
		"got the expected error for an empty username",
	)
	assert.Equal(t, User{}, user, "user returned on error is empty") // HL
}
