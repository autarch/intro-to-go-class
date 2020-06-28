package user

import "testing"

func TestNew(t *testing.T) { // HL
	user, err := NewUser("", "password")

	if err == nil {
		t.Error("expected an error when the username is empty")
	}

	empty := User{}
	if user != empty {
		t.Error("expected an empty user when the username is empty")
	}
}
