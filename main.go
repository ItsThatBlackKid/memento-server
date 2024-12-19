package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"memento/context"
	"memento/controller"
	"net/http"
	"os"
)

var DB *gorm.DB

func initialize() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	// initalize app
	initialize()

	// define routes
	r := mux.NewRouter()
	context.Context = context.RequestContext{
		DB: DB,
	}

	// user + auth routes
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", controller.GetUser).Methods("GET")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")

	// memento routes
	r.HandleFunc("/memento", controller.CreateMemento).Methods("POST")
	r.HandleFunc("/memento/{userid}", controller.GetMementos).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
