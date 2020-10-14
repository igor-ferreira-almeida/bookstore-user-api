package service

import (
	"github.com/igor-ferreira-almeida/bookstore-user-api/domain/user"
	"github.com/igor-ferreira-almeida/bookstore-user-api/util/exception"
)

func CreateUser(user user.User) (*user.User, *exception.Exception) {
	if ex:= user.Validade(); ex != nil {
		return nil, ex
	}

	if ex:= user.Save(); ex != nil {
		return nil, ex
	}

	return &user, nil
}

func GetUser(id int64) (*user.User, *exception.Exception) {
	user := &user.User{ID: id}
	if ex:= user.Get(); ex != nil {
		return nil, ex
	}
	return user, nil
}