package models

import (
	"errors"
	"fmt"
)

// User model.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers retrive the list added on the application.
func GetUsers() []*User {
	return users
}

// AddUser a the user passed by parameter to the application's user list.
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include the id or it must to be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID searchs and return the users saved on the application.
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not funded", id)
}

// UpdateUser udpates an saved user on the application.
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not funded", u.ID)
}

// RemoveUserByID delete a user on the application.
func RemoveUserByID(id int) error {
	for i, candidate := range users {
		if candidate.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not funded", id)
}
