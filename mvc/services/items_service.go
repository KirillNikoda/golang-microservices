package services

import (
	"net/http"

	"github.com/KirillNikoda/golang-microservices/mvc/domain"
	"github.com/KirillNikoda/golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId int) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
