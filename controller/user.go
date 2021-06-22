package controller

import (
	"encoding/json"
	"golearn/model"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request, c *model.Claims) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := model.GetUsers(start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithSuccess(w, http.StatusOK, c, "Success get users", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request, c *model.Claims) {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	defer r.Body.Close()

	if err := u.CreateUser(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithSuccess(w, http.StatusCreated, c, "Success create user", u)
}