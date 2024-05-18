package models

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
