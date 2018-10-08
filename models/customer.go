package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Customer model
type Customer struct {
	gorm.Model

	FirstName string     `gorm:"size:100;not null;index" binding:"required"`
	LastName  string     `gorm:"size:100;not null;index" binding:"required"`
	BirthDate *time.Time `gorm:"not null" binding:"required"`
	Gender    string     `gorm:"not null" binding:"required"`
	Email     string     `gorm:"not null" binding:"required"`
	Address   string     `gorm:"size:200"`
}
