package errors

import (
	"errors"
)

var (
	ErrTooManyAttempts        = errors.New("Password verification attempts exceeded, please wait for one mintue to retry")
	ErrIncorrectCredentials   = errors.New("Incorrect username or password")
	ErrUsernameAlreadyExist   = errors.New("Username already exists")
	ErrUsernameLengthTooShort = errors.New("Invalid Username Length: Username length should be longer than 8")
	ErrUsernameLengthTooLong  = errors.New("Invalid Username Length: Username length should be shorter than 32")
	ErrPasswordLengthTooShort = errors.New("Invalid Password Length: Password length should be longer than 8")
	ErrPasswordLengthTooLong  = errors.New("Invalid Password Length: Password length should be shorter than 32")
	ErrPasswordFormat         = errors.New("Invalid Password: Password must contain at least 1 uppercase letter, 1 lowercase letter, and 1 number")
	ErrFailedToUpdateAcInfor  = errors.New("Failed to update account failed attempts")
)

type Err struct {
	Code int
	Err  error
}
