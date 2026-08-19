package models

import "time"

type Resume struct {
	Id         int64     `json:"id,omitempty"`
	UserID     int64     `json:"user_id,omitempty"`
	Name       string    `json:"name" binding:"required,min=1,max=255"`
	File_path  string    `json:"file_path" binding:"required,max=500"`
	Created_at time.Time `json:"created_at"`
}
