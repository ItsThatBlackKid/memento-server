package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"memento/context"
	"memento/controller"
	"memento/models"
	"net/http"
	"os"
)

var DB *gorm.DB

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})
	DB = DB.Set("gorm:auto_preload", true)
	log.Println("Loaded database")

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Memento{}); err != nil {
		log.Fatal(err)
	}
	log.Println("User table migrated")
	if err := DB.AutoMigrate(&models.Memento{}); err != nil {
		log.Fatal(err)
	}
	log.Println("Memento table migrated")
}

func main() {
	// initalize app
	initDB()

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
	r.HandleFunc("/memento/{userid: [0-9]+}", controller.GetMementos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
