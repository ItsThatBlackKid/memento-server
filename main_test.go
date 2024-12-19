package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initialize()
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
	if _, err := DB.Exec(tableConnectionQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	DB.Exec("DELETE FROM memento")
	DB.Exec("DELETE FROM user")
}
