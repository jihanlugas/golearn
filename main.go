package main

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golearn/controller"
	"golearn/model"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type httpHandlerFunc func(http.ResponseWriter, *http.Request)
type httpHandlerFuncNext func(http.ResponseWriter, *http.Request, *model.Claims)

func authMiddleware(next httpHandlerFuncNext) httpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				controller.RespondWithError(w, http.StatusUnauthorized, "Unautorized")
				return
			}
			controller.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		tokenStr := cookie.Value

		claims := &model.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, err
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				controller.RespondWithError(w, http.StatusUnauthorized, "Unautorized")
				return
			}
			controller.RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < time.Second * 0 {
		//	controller.RespondWithError(w, http.StatusUnauthorized, "Unautorized")
		//	return
		//}

		if !tkn.Valid {
			controller.RespondWithError(w, http.StatusUnauthorized, "Unautorized")
			return
		}

		next(w, r, claims)
	}
}

func main() {
	log.Println("Listening server at http://localhost:8010")
	router := mux.NewRouter()

	router.HandleFunc("/signin", controller.Signin).Methods("POST")

	router.HandleFunc("/users", authMiddleware(controller.GetUsers)).Methods("GET")
	router.HandleFunc("/user", authMiddleware(controller.CreateUser)).Methods("POST")

	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8010", router))
}
