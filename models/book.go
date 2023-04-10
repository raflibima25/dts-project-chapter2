package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NameBook  string    `json:"name_book" gorm:"not null;type:varchar(191)"`
	Author    string    `json:"author" gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book before create")

	if len(b.NameBook) <= 8 {
		err = errors.New("name book is too short")
	}

	if len(b.Author) <= 5 {
		err = errors.New("author is too short")
	}

	return
}
