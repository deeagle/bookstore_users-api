package services

import (
	"myapp/domain/users"
	"myapp/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
