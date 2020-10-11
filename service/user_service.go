package service

import "github.com/igor-ferreira-almeida/bookstore-user-api/domain/user"

func CreateUser(user user.User) (user.User, error) {

	return user, nil
}