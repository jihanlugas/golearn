package controller

import (
	"encoding/json"
	"golearn/model"
	"net/http"
	"time"
)

type ResponseSuccess struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseErrorLogout struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Logout bool `json:"logout"`
}


func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := &ResponseError{
		Error: true,
		Message: message,
	}
	RespondWithJSON(w, code, response)
}

func RespondWithErrorLogout(w http.ResponseWriter, code int, message string) {
	response := &ResponseErrorLogout{
		Error: true,
		Message: message,
		Logout: true,
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      "",
		Expires: 	time.Now().Add(time.Minute * 0),
		Domain: "localhost",
		Path: "/",
	})

	RespondWithJSON(w, code, response)
}

func RespondWithSuccess(w http.ResponseWriter, code int, c *model.Claims, message string, payload interface{}) {
	token, err := model.GenerateToken(c.Email)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      token,
		Expires: 	time.Now().Add(time.Minute * 60),
		Domain: "localhost",
		Path: "/",
	})

	response := &ResponseSuccess{
		Success: true,
		Message: message,
		Data: payload,
	}

	RespondWithJSON(w, code, response)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJSONAddToken(w http.ResponseWriter, code int, message string, payload interface{}, token string) {
	res := &ResponseSuccess{
		Message: message,
		Success: true,
		Data: payload,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      token,
		Expires: 	time.Now().Add(time.Minute * 60),
		Domain: "localhost",
		Path: "/",
	})

	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJSONRemoveToken(w http.ResponseWriter, code int, message string) {
	response := &ResponseSuccess{
		Success: true,
		Message: message,
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      "",
		Expires: 	time.Now().Add(time.Minute * 0),
		Domain: "localhost",
		Path: "/",
	})

	RespondWithJSON(w, code, response)

}
