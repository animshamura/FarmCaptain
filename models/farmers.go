package models

import "github.com/jinzhu/gorm"

// Farmer represents a farmer entity
type Farmer struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
