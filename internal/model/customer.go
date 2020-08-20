package model

import "time"

type Customer struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"deleted_at"`
	Goal              string     `json:"goal"`
	Allergy           string     `json:"allergy"`
	Weight            float64    `json:"weight"`
	Height            string     `json:"height"`
	DietaryPreference string     `json:"dietary_preference"`
	Gender            string     `json:"gender"`
	UserID            uint       `json:"user_id"`
	User              *User      `json:"user,omitempty"`
	DietitianID       uint       `json:"dietitian_id"`
	Dietitian         *Dietitian `gorm:"foreignKey:DietitianID" json:"dietitian"`
	MealPlan          *MealPlan  `json:"meal_plan"`
}

func (c Customer) TableName() string {
	return "customers"
}
