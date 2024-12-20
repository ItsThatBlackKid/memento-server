package models

import (
	"gorm.io/gorm"
	"log"
	"memento/context"
)

type Memento struct {
	gorm.Model
	User   User   `gorm:"references:ID" json:"user,omitempty"`
	UserID int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Mood   int8   `gorm:"check:mood <=10" json:"mood"`
}

func (m *Memento) GetMemento(db *gorm.DB) error {
	result := db.First(&m, m.ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Memento) CreateMemento() error {
	result := context.Context.DB.Preload("User").Create(&m)

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
	result := db.Find(&mementos, "userid=$1", m.UserID)

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
