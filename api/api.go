package api

import (
	"context"
	"net/http"
	"user_api/models"
	"user_api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	collect *mongo.Collection
}

func NewApi(collect *mongo.Collection) *Api {
	return &Api{
		collect: collect,
	}
}

// CreateAccountHandler godoc
// @Summary      Create an account
// @Description  Create an account by the desired username and password.
// @Description  Enter the username and password,
// @Description
// @Description  The username must meet the following criteria:
// @Description  - Minimum length of 3 characters and a maximum length of 32 characters.
// @Description
// @Description  The password must meet the following criteria:
// @Description  - Minimum length of 8 characters and maximum length of 32 characters.
// @Description  - Must contain at least 1 uppercase letter, 1 lowercase letter, and 1 number.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account	 	body      models.AccountRequest 					true	"Account credentials"
// @Success      200  			{object}  models.ResponseType.Response
// @Failure      400  			{object}  models.ResponseType.BadRequestResponse	"Invalid username or password"
// @Failure      409  			{object}  models.ResponseType.BadRequestResponse	"Account already exists"
// @Router       /v1/signup [post]
func (a *Api) CreateAccountHandler(c *gin.Context) {
	var req models.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			Success: false,
			Reason:  err.Error(),
		})
		return
	}

	if err := services.CreateAccount(context.Background(), a.collect, &models.Account{
		Username:       req.Username,
		Password:       req.Password,
		FailedAttempts: 0,
	}); err != nil {
		c.JSON(err.Code, &models.Response{
			Success: false,
			Reason:  err.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &models.Response{
		Success: true,
	})
}

// VerifyAccountHandler godoc
// @Summary      Verify an account
// @Description  Verifies the provided account credentials.
// @Description  If the password verification fails five times, the user is required to wait for one minute before attempting again.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account 		body      models.AccountRequest 					true	"Account credentials"
// @Success      200 			{object}  models.ResponseType.Response
// @Failure      401  			{object}  models.ResponseType.BadRequestResponse	"Incorrect username or password"
// @Failure      429  			{object}  models.ResponseType.BadRequestResponse	"Password attempts exceed "
// @Router       /v1/login [post]
func (a *Api) VerifyAccountHandler(c *gin.Context) {
	var req models.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			Success: false,
			Reason:  err.Error(),
		})
		return
	}

	ac, err := services.VerifyAccount(context.Background(), a.collect, &models.Account{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(err.Code, &models.Response{
			Success: false,
			Reason:  err.Err.Error(),
		})
		return
	}

	if _, err := ac.UpdateAccountFailedAttempts(context.Background(), a.collect, 0 /* Reset */); err != nil {
		c.JSON(err.Code, &models.Response{
			Success: false,
			Reason:  err.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.Response{
		Success: true,
	})
}
