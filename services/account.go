package services

import (
	"context"
	"net/http"
	"user_api/errors"
	"user_api/models"
	"user_api/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateAccount checks if the format of the input account valid,
// then hash the password and insert to DB.
func CreateAccount(ctx context.Context, coll *mongo.Collection, a *models.Account) *errors.Err {
	if err := a.ValidAccount(); err != nil {
		return err
	}

	// If the username of the input account can be found from DB, that means the username has already be registered.
	if _, err := a.FindAccountFromDB(ctx, coll); err == nil {
		return &errors.Err{
			Code: http.StatusConflict,
			Err:  errors.ErrUsernameAlreadyExist,
		}
	}

	var err error
	a.Password, err = utils.HashPassword(a.Password)
	if err != nil {
		return &errors.Err{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	_, err = coll.InsertOne(ctx, a)
	if err != nil {
		return &errors.Err{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return nil
}

// VerifyAccount verifies the format of the input account is valid, confirm the input username existed.
// Then checks the verification attempt is less than 5 times.
// Finally, unhash the password stored in DB and compares to the input password.
func VerifyAccount(ctx context.Context, coll *mongo.Collection, a *models.Account) (*models.Account, *errors.Err) {
	res, err := a.FindAccountFromDB(ctx, coll)
	if err != nil {
		return nil, err
	}

	if err := res.CheckVerificationAttempts(ctx, coll); err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(a.Password, res.Password); err != nil {
		res.UpdateAccountFailedAttempts(ctx, coll, res.FailedAttempts+1)
		return nil, &errors.Err{
			Code: http.StatusUnauthorized,
			Err:  errors.ErrIncorrectCredentials,
		}
	}

	return res, nil
}
