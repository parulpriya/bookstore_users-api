package users

import (
	"fmt"
	"github.com/parulpriya/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestErr(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.BadRequestErr(fmt.Sprintf("user %d already exists", user.ID))
	}

	usersDB[user.ID] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NotFoundErr(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}
