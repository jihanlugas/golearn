package controller

import (
	"database/sql"
	"encoding/json"
	"golearn/model"
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

	//RespondWithError(w, http.StatusNotFound, "Invalid Email or Password")
	//return

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
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSONAddToken(w, http.StatusOK, "Login Success", map[string]string{"result": "success"}, token)
	return
}

func Signout(w http.ResponseWriter, r *http.Request) {
	RespondWithJSONRemoveToken(w, http.StatusOK, "Success Logout")
	return
}