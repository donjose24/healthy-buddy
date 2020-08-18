package model

import "github.com/jinzhu/gorm"

type Customer struct {
	*gorm.Model
	Goal              string    `json:"goal"`
	Allergy           string    `json:"allergy"`
	Weight            float64   `json:"weight"`
	Height            string    `json:"height"`
	DietaryPreference string    `json:"dietary_preference"`
	Gender            string    `json:"gender"`
	UserID            uint      `json:"user_id"`
	DietitianID       uint      `json:"dietitian_id"`
	Dietitian         Dietitian `json:"dietitian"`
}

func (c Customer) TableName() string {
	return "customers"
}
