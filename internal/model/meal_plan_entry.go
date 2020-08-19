package model

import "time"

type MealPlanEntry struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	MealPlanID  uint       `json:"-"`
	MealPlan    *MealPlan  `json:"meal_plan,omitempty"`
	MealTime    string     `json:"meal_time"`
	Food        string     `json:"food"`
	Protein     float64    `json:"protein"`
	Grams       float64    `json:"grams"`
	Fat         float64    `json:"fat"`
	Carb        float64    `json:"carb"`
	Calories    float64    `json:"calories"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
}
