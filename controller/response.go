package controller

import (
	"encoding/json"
	"net/http"
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

func RespondWithSuccess(w http.ResponseWriter, code int, message string, payload interface{}) {
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

//func RespondWithJSONAddToken(w http.ResponseWriter, code int, payload interface{}, token string) {
//	res := &Response{
//		Success: true,
//		Data: payload,
//	}
//
//	response, _ := json.Marshal(res)
//
//	w.Header().Set("Accept", "application/json")
//	w.Header().Set("Content-Type", "application/json")
//
//	http.SetCookie(w, &http.Cookie{
//		Name:       "Authorization",
//		Value:      token,
//		Expires: 	time.Now().Add(time.Minute * 5),
//		//Domain: "tess.com",
//	})
//
//	w.WriteHeader(code)
//	w.Write(response)
//}
