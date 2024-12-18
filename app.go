package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(db string) {
	var err error
	a.DB, err = sql.Open("sqlite3", db)
	if err != nil {
		log.Fatal(err)
	}

	// read the sql file
	filePath := filepath.Join("../", "db", "mementodb.sql")
	c, ioErr := os.ReadFile(filePath)
	if ioErr != nil {
		log.Fatal(ioErr)
	}

	// execute table creation query
	tableCreationQuery := string(c)
	_, err = a.DB.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
