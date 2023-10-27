package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
	UserId      string `json:"user_id"`
	Jobs        []Job  `json:"jobs" gorm:"foreignKey:CompanyID"`
}

type NewCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	FoundedYear string `json:"founded_year" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Jobs        []Job  `json:"jobs"`
}

type Job struct {
	gorm.Model
	Title           string `json:"title"`
	ExperienceLevel string `json:"experience_required"`
	CompanyID       uint   `json:"company_id"`
}

/*
{
    "company_name": "ABC Inc.",
    "founded_year": "2000",
    "location": "New York",
    "user_id": "12345"
}
*/