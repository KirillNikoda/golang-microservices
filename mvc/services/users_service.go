package services

import (
	"github.com/KirillNikoda/golang-microservices/mvc/domain"
	"github.com/KirillNikoda/golang-microservices/mvc/utils"
)

func GetUser(userId int) (*domain.User, *utils.ApplicationError) {
	user, err := domain.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
