package model

import "time"

type MealPlan struct {
	ID              uint            `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *time.Time      `sql:"index" json:"deleted_at"`
	DietitianID     uint            `json:"-"`
	Dietitian       *Dietitian      `json:"dietitian,omitempty"`
	CustomerID      uint            `json:"-"`
	Customer        *Customer       `json:"customer,omitempty"`
	StartDate       string          `json:"start_date"`
	EndDate         string          `json:"end_date"`
	Remarks         string          `json:"remarks"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	MealPlanEntries []MealPlanEntry `json:"entries,omitempty"`
}
