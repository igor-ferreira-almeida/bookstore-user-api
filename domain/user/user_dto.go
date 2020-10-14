package user

import (
	"github.com/igor-ferreira-almeida/bookstore-user-api/util/exception"
	"strings"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func NewUser(id int64, firstName, LastName, email, dateCreated string) User {
	return User{
		ID:          id,
		FirstName:   firstName,
		LastName:    LastName,
		Email:       email,
		DateCreated: dateCreated,
	}
}

func (user *User) Validade() *exception.Exception {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return exception.NewBadRequestException("Invalid e-mail address")
	}
	return nil
}