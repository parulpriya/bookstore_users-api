package services

import (
	"github.com/parulpriya/bookstore_users-api/domain/users"
	"github.com/parulpriya/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	var user users.User
	user.ID = userID

	err := user.Get()
	if err != nil {
		return nil, err
	}
	return &user, nil
}