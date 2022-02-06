package services

import (
	"myapp/domain/users"
	"myapp/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	userErr := user.Validate()
	if userErr != nil {
		return nil, userErr
	}

	saveErr := user.Save()
	if saveErr != nil {
		return nil, saveErr
	}

	return &user, nil
}
