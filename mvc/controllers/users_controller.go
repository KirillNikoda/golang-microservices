package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KirillNikoda/golang-microservices/mvc/services"
)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("user_id must be a number"))
		return
	}

	user, err := services.GetUser(int(userId))
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	rw.Write(jsonValue)
}
