package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"memento/appContext"
	"memento/controller"
	"memento/middleware"
	"memento/models"
	"net/http"
	"os"
)

func initDB() {
	var DB, err = gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})
	DB = appContext.DB.Set("gorm:auto_preload", true)
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

	appContext.DB = DB
}

func main() {
	// ensure .env file is loaded
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	initDB()

	r := mux.NewRouter()
	// user + auth routes
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")

	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)
	authRouter.HandleFunc("/users/{id:[0-9]+}", controller.GetUser).Methods("GET")

	// memento routes
	authRouter.HandleFunc("/memento", controller.CreateMemento).Methods("POST")
	authRouter.HandleFunc("/memento/{userid: [0-9]+}", controller.GetMementos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
