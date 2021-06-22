package controller

import (
	"database/sql"
	"encoding/json"
	"golearn/model"
	"log"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var c model.Credentials
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.Signin(); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Invalid Email or Password")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	token, err := model.GenerateToken(c.Email)
	if err != nil {
		log.Println("asd")
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSONAddToken(w, http.StatusOK, map[string]string{"result": "success"}, token)
}