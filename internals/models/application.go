package models

import "time"

type Application struct {
	Id              int64     `json:"id"`
	User_id         int64     `json:"user_id,omitempty"`
	Company_name    string    `json:"company_name" binding:"required,min=1,max=255"`
	Job_title       string    `json:"job_title" binding:"required,min=1,max=255"`
	Job_url         string    `json:"job_url" binding:"required"`
	Location        string    `json:"location" binding:"required,min=1,max=255"`
	Employment_type string    `json:"employment_type"`
	Salary_min      int       `json:"salary_min"`
	Salary_max      int       `json:"salary_max"`
	Status          string    `json:"status" binding:"required"`
	Applied_at      time.Time `json:"applied_at"`
	Notes           string    `json:"notes" binding:"max=1000"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}

type ApplicationResponse struct {
	Company_name    string    `json:"company_name" binding:"required"`
	Job_title       string    `json:"job_title" binding:"required"`
	Job_url         string    `json:"job_url" binding:"required"`
	Location        string    `json:"location" binding:"required"`
	Employment_type string    `json:"employment_type"`
	Salary_min      int       `json:"salary_min"`
	Salary_max      int       `json:"salary_max"`
	Status          string    `json:"status" binding:"required"`
	Applied_at      time.Time `json:"applied_at"`
	Notes           string    `json:"notes"`
}
