package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Password  string `json:"-"`
}
