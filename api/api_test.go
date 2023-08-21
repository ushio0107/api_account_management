package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user_api/errors"
	"user_api/models"
	"user_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestSignin(t *testing.T) {
	existUsername := utils.RandomStringWithAllChar(3, 32)
	for _, testcase := range []struct {
		testName       string
		body           gin.H
		expectedCode   int
		verifyResponse func(*models.Response)
	}{
		{
			testName: "create_valid_account",
			body: gin.H{
				"username": existUsername,
				"password": utils.RandomStringWithAllChar(8, 32),
			},
			expectedCode: http.StatusCreated,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, true)
				require.Equal(t, res.Reason, "")
			},
		}, {
			testName: "create_invalid_account_invalid_account",
			body: gin.H{
				"username": utils.RandomStringWithAllChar(8, 32)[:2],
				"password": utils.RandomStringWithAllChar(8, 32),
			},
			expectedCode: http.StatusBadRequest,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrUsernameLengthTooShort.Error())
			},
		}, {
			testName: "create_invalid_account_existed_account",
			body: gin.H{
				"username": existUsername,
				"password": utils.RandomStringWithAllChar(8, 32),
			},
			expectedCode: http.StatusConflict,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrUsernameAlreadyExist.Error())
			},
		}, {
			testName: "create_invalid_account_invalid_password_length_less",
			body: gin.H{
				"username": utils.RandomStringWithAllChar(3, 32),
				"password": utils.RandomStringWithAllChar(1, 7),
			},
			expectedCode: http.StatusBadRequest,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrPasswordLengthTooShort.Error())
			},
		}, {
			testName: "create_invalid_account_invalid_password_length_more",
			body: gin.H{
				"username": utils.RandomStringWithAllChar(3, 32),
				"password": utils.RandomStringWithAllChar(33, 50),
			},
			expectedCode: http.StatusBadRequest,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrPasswordLengthTooLong.Error())
			},
		}, {
			testName: "create_invalid_account_invalid_password_format",
			body: gin.H{
				"username": utils.RandomStringWithAllChar(3, 32),
				"password": utils.RandomStringWithAlp(8, 32),
			},
			expectedCode: http.StatusBadRequest,
			verifyResponse: func(res *models.Response) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrPasswordFormat.Error())
			},
		},
	} {
		input, err := json.Marshal(testcase.body)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/v1/signup", bytes.NewReader(input))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		require.Equal(t, testcase.expectedCode, w.Code)

		var response models.Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		testcase.verifyResponse(&response)
	}

}

func TestLogin(t *testing.T) {
	testUsername := utils.RandomStringWithAllChar(3, 32)
	testPassword := utils.RandomStringWithAllChar(8, 32)
	body := gin.H{
		"username": testUsername,
		"password": testPassword,
	}
	input, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/signup", bytes.NewReader(input))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response models.Response
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	for _, testcase := range []struct {
		testName       string
		body           gin.H
		attempTimes    int
		expectedCode   int
		verifyResponse func(*models.Response, int)
	}{
		{
			testName: "login_success",
			body: gin.H{
				"username": testUsername,
				"password": testPassword,
			},
			attempTimes:  1,
			expectedCode: http.StatusOK,
			verifyResponse: func(res *models.Response, _ int) {
				require.Equal(t, res.Success, true)
				require.Equal(t, res.Reason, "")
			},
		}, {
			testName: "login_failed_wrong_username",
			body: gin.H{
				"username": "unexistedusername",
				"password": testPassword,
			},
			attempTimes:  1,
			expectedCode: http.StatusUnauthorized,
			verifyResponse: func(res *models.Response, _ int) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrIncorrectCredentials.Error())
			},
		}, {
			testName: "login_failed_wrong_password",
			body: gin.H{
				"username": testUsername,
				"password": utils.RandomStringWithAllChar(8, 32),
			},
			attempTimes:  1,
			expectedCode: http.StatusUnauthorized,
			verifyResponse: func(res *models.Response, _ int) {
				require.Equal(t, res.Success, false)
				require.Equal(t, res.Reason, errors.ErrIncorrectCredentials.Error())
			},
		}, {
			testName: "login_failed_wrong_password_5_times",
			body: gin.H{
				"username": testUsername,
				"password": utils.RandomStringWithAllChar(8, 32),
			},
			attempTimes:  7,
			expectedCode: http.StatusUnauthorized,
			verifyResponse: func(res *models.Response, attempTimes int) {
				require.Equal(t, res.Success, false)
				if attempTimes < 4 { // Have logged in to same account for 1 time, so 4 times left.
					require.Equal(t, res.Reason, errors.ErrIncorrectCredentials.Error())
				} else {
					require.Equal(t, res.Reason, errors.ErrTooManyAttempts.Error())
				}
			},
		},
	} {
		for i := 0; i < testcase.attempTimes; i++ {
			input, err := json.Marshal(testcase.body)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/v1/login", bytes.NewReader(input))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if i < 4 { // Have logged in to same account for 1 time, so 4 times left.
				require.Equal(t, testcase.expectedCode, w.Code)
			} else {
				require.Equal(t, http.StatusTooManyRequests, w.Code)
			}

			var response models.Response
			err = json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			testcase.verifyResponse(&response, i)
		}
	}
}
