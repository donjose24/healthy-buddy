package model

import "github.com/jinzhu/gorm"

type Dietitian struct {
	*gorm.Model
	Specialty         string     `json:"specialty"`
	YearsOfExperience int        `json:"years_of_experience"`
	Customers         []Customer `json:"customers"`
}
