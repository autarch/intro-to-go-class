package main

import "fmt"

func NewUser(id int) (User, error) {
	user := lookupUser(id)
	if user == nil {
		return nil, fmt.Errorf("No user matching id %d was found", id)
	}

	return user, nil
}
