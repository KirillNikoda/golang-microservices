package services

import (
	"net/http"
	"testing"

	"github.com/KirillNikoda/golang-microservices/mvc/domain"
	"github.com/KirillNikoda/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var (
	UsersDaoMock    usersDaoMock
	getUserFunction func(userId int) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exist",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
