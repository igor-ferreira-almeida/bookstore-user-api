package user

import (
	"fmt"
	"github.com/igor-ferreira-almeida/bookstore-user-api/util/exception"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *exception.Exception {
	result := usersDB[user.ID]

	if result == nil {
		return exception.NewNotFoundException(fmt.Sprintf("User %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *exception.Exception {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return exception.NewBadRequestException(fmt.Sprintf("e-mail %s already registered", user.Email))
		}
		return exception.NewBadRequestException(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
