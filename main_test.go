package main

import (
	"log"
	"os"
	"testing"
)

var app = App{}

func TestMain(m *testing.M) {
	app.Initialize(os.Getenv("DB"))
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

const tableConnectionQuery = `
	CREATE TABLE IF NOT EXISTS  [user] (
    id integer primary key  not null,
    username text not null,
    email text not null,
    first_name text,
    last_name text,
    password text not null
);

CREATE TABLE IF NOT EXISTS [memento] (
    id integer primary key not null,
    userid integer not null,
    title text not null,
    body text,
    foreign key (userid) references user(id) on update cascade
)
`

func ensureTableExists() {
	if _, err := app.DB.Exec(tableConnectionQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM memento")
	app.DB.Exec("DELETE FROM user")
}
