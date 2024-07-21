package models

type CreateRoom struct {
	RoomName     string   `json:"room_name"`
	RoomType     string   `json:"room_type"`
	FirstMessage string   `json:"first_message"`
	UserIds      []string `json:"user_ids"`
}
