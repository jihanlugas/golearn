package controller

import (
	"encoding/json"
	"golearn/model"
	"log"
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

func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := &ResponseError{
		Error: true,
		Message: message,
	}
	RespondWithJSON(w, code, response)
}

func RespondWithSuccess(w http.ResponseWriter, code int, c *model.Claims, message string, payload interface{}) {
	token, err := model.GenerateToken(c.Email)
	if err != nil {
		log.Println("asd")
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      token,
		Expires: 	time.Now().Add(time.Minute * 1),
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

func RespondWithJSONAddToken(w http.ResponseWriter, code int, payload interface{}, token string) {
	res := &ResponseSuccess{
		Success: true,
		Data: payload,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      token,
		Expires: 	time.Now().Add(time.Minute * 1),
		Domain: "localhost",
		Path: "/",
	})

	w.WriteHeader(code)
	w.Write(response)
}
