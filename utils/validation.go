package utils

import (
	"unicode"
	"user_api/errors"
)

func ValidUsername(username string) error {
	if len(username) < 3 {
		return errors.ErrUsernameLengthTooShort
	}
	if len(username) > 32 {
		return errors.ErrUsernameLengthTooLong
	}

	// Users can also add some instructions to verify if a username valid.

	return nil
}

func ValidPassword(password string) error {
	if len(password) < 8 {
		return errors.ErrPasswordLengthTooShort
	}
	if len(password) > 32 {
		return errors.ErrPasswordLengthTooLong
	}

	hasUpper := false
	hasLower := false
	hasNumber := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasNumber = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber {
		return errors.ErrPasswordFormat
	}

	return nil
}
