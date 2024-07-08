package models

import "time"

type Message struct {
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	ID          string    `json:"id" db:"id"`
	SenderId    string    `json:"sender_id" db:"sender_id"`
	Message     string    `json:"message" db:"message"`
	FullName    string    `json:"full_name" db:"full_name"`
}
