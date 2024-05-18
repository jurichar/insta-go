package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uint   `json:"id" gorm:"primaryKey"`
	Email             string `json:"email" gorm:"unique"`
	Password          string `json:"password" gorm:"-"` // - means that this field will not be stored in the database
	CreatedAt         time.Time
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ConfirmationToken string `json:"confirmation_token"`
	Confirmed         bool   `json:"confirmed"`
}
