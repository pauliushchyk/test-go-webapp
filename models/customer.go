package models

import (
	"time"
)

// Customer model
type Customer struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `form:"-" binding:"-"`
	UpdatedAt time.Time  `form:"-" binding:"-"`
	DeletedAt *time.Time `form:"-" sql:"index" binding:"-"`

	FirstName string    `form:"first_name" gorm:"size:100;not null;index" json:"first_name" binding:"required,max=100"`
	LastName  string    `form:"last_name" gorm:"size:100;not null;index" json:"last_name" binding:"required,max=100"`
	BirthDate time.Time `form:"birth_date" gorm:"not null" json:"birth_date" binding:"exists" time_format:"2006-01-02"`
	Gender    string    `form:"gender" gorm:"not null" json:"gender" binding:"required,eq=female|eq=male"`
	Email     string    `form:"email" gorm:"not null" json:"email" binding:"required,email"`
	Address   string    `form:"address" gorm:"size:200" json:"address" binding:"max=200"`
}
