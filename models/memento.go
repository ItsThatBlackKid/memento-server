package models

import (
	"database/sql"
	"errors"
)

type Memento struct {
	ID     int16  `json:"id"`
	userid int16  `json:"userid"`
	title  string `json:"title"`
	body   string `json:"body"`
}

func (m *Memento) getMemento(db *sql.DB) error {
	return errors.New("not implemented")
}

func (m *Memento) createMemento(db *sql.DB) error {
	return errors.New("not implemented")
}

func (m *Memento) deleteMemento(db *sql.DB) error {
	return errors.New("not implemented")
}

func (m *Memento) getMementos(db *sql.DB) ([]Memento, error) {
	return nil, errors.New("not implemented")
}
