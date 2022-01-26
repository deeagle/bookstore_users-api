package services

import (
	"myapp/domain/users"
	"myapp/utils/errors"
	"strings"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, errors.NewBadRequestError("invalid email address")
	}
	return &user, nil
}
