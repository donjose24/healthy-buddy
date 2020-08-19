package model

import "time"

type MealPlanEntry struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	MealPlanID  uint       `json:"-"`
	MealPlan    MealPlan   `json:"meal_plan"`
	MealTime    string     `json:"meal_time"`
	Food        string     `json:"food"`
	Protein     float64    `json:"protein"`
	Fat         float64    `json:"fat"`
	Carb        float64    `json:"carb"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
}
