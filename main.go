package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"memento/controller"
	"net/http"
	"os"
	"path/filepath"
)

var DB *sql.DB

func initialize() {
	var err error
	DB, err = sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		log.Fatal(err)
	}

	// read the sql file
	filePath := filepath.Join("./", "db", "mementodb.sql")
	c, ioErr := os.ReadFile(filePath)
	if ioErr != nil {
		log.Fatal(ioErr)
	}

	// execute table creation query
	tableCreationQuery := string(c)
	_, err = DB.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// initalize app
	initialize()

	// define routes
	r := mux.NewRouter()
	uc := controller.UserController{
		DB: DB,
	}

	r.HandleFunc("/users", uc.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", uc.GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
