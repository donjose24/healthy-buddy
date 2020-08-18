package model

import "time"

type Dietitian struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"deleted_at"`
	Specialty         string     `json:"specialty"`
	YearsOfExperience int        `json:"years_of_experience"`
	Customers         []Customer `json:"customers"`
	UserID            uint       `json:"-"`
	User              User       `gorm:"foreignKey:UserID" json:"user"`
}
