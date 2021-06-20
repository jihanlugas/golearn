package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"golearn/controller"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening server at http://localhost:8010")
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8010", router))
}
