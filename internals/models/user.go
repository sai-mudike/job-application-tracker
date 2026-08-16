package models

import "time"

type User struct {
	Id         int64     `json:"id"`
	UserName   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
