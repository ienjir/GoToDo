package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User2 struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	UserName  string    `gorm:"unique;size:40;not null"`
	Email     string    `gorm:"unique;size:150;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ToDos     []ToDo    `gorm:"foreignKey:UserID"`
}

type ToDo struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:512"`
	Location    string    `gorm:"size:100"`
	DueDate     time.Time `gorm:""`
}
