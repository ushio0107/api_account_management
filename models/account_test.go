package models

import (
	"testing"
	"user_api/errors"
	"user_api/utils"

	"github.com/stretchr/testify/require"
)

const (
	min = 8
	max = 32
)

func RandomValidUser() string {
	return utils.RandomStringWithAllChar(min, max)
}

func RandomInvalidUserLessThan3() string {
	return utils.RandomStringWithAllChar(min, max)[:2]
}

func RandomInvalidUserMoreThan32() string {
	return utils.RandomStringWithAllChar(max+1, 50)
}

func RandomValidPassword() string {
	return utils.RandomStringWithAllChar(min, max)
}

func RandomInvalidPasswordFormat() string {
	return utils.RandomStringWithAlp(min, max)
}

func RandomInvalidPasswordLessThan8() string {
	return utils.RandomStringWithAllChar(1, min-1)
}

func RandomInvalidPasswordMoreThan32() string {
	return utils.RandomStringWithAllChar(max+1, 50)
}

func TestValidAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		account := &Account{
			Username: RandomValidUser(),
			Password: RandomValidPassword(),
		}
		err := account.ValidAccount()

		require.Empty(t, err)
	}
}

func TestInvalidAccount(t *testing.T) {
	for _, f := range []func() string{
		RandomInvalidUserLessThan3,
		RandomInvalidUserMoreThan32,
	} {
		account := &Account{
			Username: f(),
			Password: RandomValidPassword(),
		}
		err := account.ValidAccount()
		require.Error(t, err.Err)
		if len(account.Username) < 3 {
			require.ErrorIs(t, err.Err, errors.ErrUsernameLengthTooShort)
		} else {
			require.ErrorIs(t, err.Err, errors.ErrUsernameLengthTooLong)
		}

	}

	for i := 0; i < 10; i++ {
		for _, f := range []func() string{
			RandomInvalidPasswordFormat,
			RandomInvalidPasswordLessThan8,
			RandomInvalidPasswordMoreThan32,
		} {
			account := &Account{
				Username: RandomValidUser(),
				Password: f(),
			}
			err := account.ValidAccount()
			require.Error(t, err.Err)
		}
	}
}
