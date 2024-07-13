package models

type CreateRoom struct {
	RoomName string   `json:"room_name"`
	UserIds  []string `json:"user_ids"`
}
