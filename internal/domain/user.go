package domain

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidID      = errors.New("invalid user id")
	ErrDivisionByZero = errors.New("division by zero")
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
