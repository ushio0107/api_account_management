package models

import (
	"context"
	"net/http"
	"time"
	"user_api/errors"
	"user_api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	Username          string    `bson:"username"`
	Password          string    `bson:"password"`
	FailedAttempts    int       `bson:"failed_attempts"`
	LastFailedAttempt time.Time `bson:"last_failed_attempt"`
}

// ValidAccount checks both username and password of the account are valid.
func (a *Account) ValidAccount() *errors.Err {
	if err := utils.ValidUsername(a.Username); err != nil {
		return &errors.Err{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	if err := utils.ValidPassword(a.Password); err != nil {
		return &errors.Err{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	return nil
}

// CheckVerificationAttempts checks the password verification attempts.
func (a *Account) CheckVerificationAttempts(ctx context.Context, coll *mongo.Collection) *errors.Err {
	if a.FailedAttempts >= 5 {
		if time.Since(a.LastFailedAttempt) < time.Minute {
			return &errors.Err{
				Code: http.StatusTooManyRequests,
				Err:  errors.ErrTooManyAttempts,
			}
		} else {
			if _, err := a.UpdateAccountFailedAttempts(ctx, coll, 0 /* Reset */); err != nil {
				return err
			}
		}
	}

	return nil
}

// FindAccountFromDB uses the username as the filter, find the specific account from DB.
func (a *Account) FindAccountFromDB(ctx context.Context, coll *mongo.Collection) (*Account, *errors.Err) {
	var res Account
	filter := bson.M{"username": a.Username}
	if err := coll.FindOne(ctx, filter).Decode(&res); err != nil /* Account can't be found */ {
		return nil, &errors.Err{
			Code: http.StatusUnauthorized,
			Err:  errors.ErrIncorrectCredentials,
		}
	}

	return &res, nil
}

// UpdateAccountFailedAttempts updates the number of the failed attempts with specific account.
func (a *Account) UpdateAccountFailedAttempts(ctx context.Context, coll *mongo.Collection, failedAttempts int) (*Account, *errors.Err) {
	a.FailedAttempts = failedAttempts
	if failedAttempts != 0 {
		a.LastFailedAttempt = time.Now()
	}
	res, err := a.UpdateAccountInfor(ctx, coll)
	if err != nil {
		return nil, &errors.Err{
			Code: http.StatusBadRequest,
			Err:  errors.ErrFailedToUpdateAcInfor,
		}
	}

	return res, nil
}

// UpdateAccountFailedAttempts updates the information of the specific account.
func (a *Account) UpdateAccountInfor(ctx context.Context, coll *mongo.Collection) (*Account, *errors.Err) {
	filter := bson.M{"username": a.Username}
	update := bson.M{"$set": bson.M{
		"failed_attempts":     a.FailedAttempts,
		"last_failed_attempt": a.LastFailedAttempt,
	}}

	var res Account
	if err := coll.FindOneAndUpdate(ctx, filter, update).Decode(&res); err != nil {
		return nil, &errors.Err{
			Code: http.StatusBadRequest,
			Err:  errors.ErrFailedToUpdateAcInfor,
		}
	}

	return &res, nil
}
