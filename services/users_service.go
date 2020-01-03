package services

import (
	"github.com/w1nsun/bookstore_users-api/domain/users"
	"github.com/w1nsun/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return nil, &errors.RestErr{
		Message: "invalid json body",
		Status:  http.StatusInternalServerError,
		Error:   "bad_request",
	}
}
