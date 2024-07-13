package models

import "time"

type Room struct {
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	UpdatedTime time.Time `json:"updated_time" db:"updated_time"`
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
}
