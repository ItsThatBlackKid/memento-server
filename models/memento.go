package models

import (
	"gorm.io/gorm"
	"log"
	"memento/context"
)

type Memento struct {
	ID     int16  `json:"id"`
	Userid int16  `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Mood   int8   `json:"mood"`
}

func (m *Memento) GetMemento(db *gorm.DB) error {
	result := db.First(&m, m.ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Memento) CreateMemento() error {
	result := context.Context.DB.Create(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Memento) DeleteMemento() error {
	result := context.Context.DB.Delete(&m)
	return result.Error
}

func (m *Memento) GetMementosByUserId(db *gorm.DB) ([]Memento, error) {
	var mementos []Memento
	result := db.Find(&mementos, "userid=$1", m.Userid)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return mementos, nil
}

func (m *Memento) GetMementos(db *gorm.DB) ([]Memento, error) {
	var mementos []Memento
	result := db.Find(&mementos)

	if result.Error != nil {
		return nil, result.Error
	}

	return mementos, nil
}
