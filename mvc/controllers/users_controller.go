package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KirillNikoda/golang-microservices/mvc/services"
	"github.com/KirillNikoda/golang-microservices/mvc/utils"
)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		rw.WriteHeader(http.StatusBadRequest)
		jsonErr, _ := json.Marshal(apiErr)
		rw.Write(jsonErr)
		return
	}

	user, apiErr := services.GetUser(int(userId))
	if apiErr != nil {
		rw.WriteHeader(apiErr.StatusCode)
		apiErr, _ := json.Marshal(apiErr)
		rw.Write(apiErr)
		return
	}

	jsonValue, _ := json.Marshal(user)
	rw.Write(jsonValue)
}
