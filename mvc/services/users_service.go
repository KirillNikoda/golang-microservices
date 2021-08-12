package services

import "github.com/KirillNikoda/golang-microservices/mvc/domain"

func GetUser(userId int) (*domain.User, error) {
	return domain.GetUser(userId)
}
