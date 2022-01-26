package services

import (
	"myapp/domain/users"
	"myapp/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	userErr := user.Validate()
	if userErr != nil {
		return nil, userErr
	}

	return &user, nil
}
