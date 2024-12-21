package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"memento/appContext"
)

type Memento struct {
	GormModel
	User   User   `json:"user"`
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Mood   int8   `gorm:"check:mood <=10" json:"mood"`
}

// Memento model hooks

func (m *Memento) AfterSave(tx *gorm.DB) (err error) {
	tx.Omit("Password").First(&(m).User, "ID", m.UserID)
	return
}

func (m *Memento) GetMemento(db *gorm.DB) error {
	result := db.First(&m, m.ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Memento) CreateMemento() error {
	result := appContext.DB.Omit("User").Create(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Memento) DeleteMemento() error {
	result := appContext.DB.Delete(&m)
	return result.Error
}

func (m *Memento) GetMementosByUserId(db *gorm.DB) ([]Memento, error) {
	var mementos []Memento
	result := db.Preload(clause.Associations).Find(&mementos, "user_id", m.UserID)

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
