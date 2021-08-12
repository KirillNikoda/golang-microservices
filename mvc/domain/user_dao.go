package domain

import (
	"fmt"
)

var (
	users = map[int]*User{
		123: {Id: 1, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}
)

func GetUser(userId int) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil

	}

	return nil, fmt.Errorf("user with id %v was not found", userId)
}
