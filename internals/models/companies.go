package models

import (
	"time"
)

type Company struct {
	Id         int64     `json:"id"`
	User_id    int       `json:"user_id,omitempty"`
	Name       string    `json:"name" binding:"required"`
	Website    string    `json:"website" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Industry   string    `json:"industry" binding:"required"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

type CompanySummary struct {
	Id   int64
	Name string
}
