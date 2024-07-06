package models

import "time"

type Message struct {
	ID          string    `json:"id"`
	SenderId    string    `json:"sender_id"`
	Message     string    `json:"message"`
	FullName    string    `json:"full_name"`
	CreatedTime time.Time `json:"created_time"`
}
