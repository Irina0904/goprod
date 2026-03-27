package main

import "errors"

var CreateUserError = errors.New("failed to create user")
var FindUserError = errors.New("user not found")

type JsonError struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"status_code"`
}
