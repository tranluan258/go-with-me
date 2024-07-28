package models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type User struct {
	Avatar   *string `json:"avatar" db:"avatar"`
	ID       string  `json:"id" db:"id"`
	Username string  `json:"username" db:"username"`
	Password string  `json:"password" db:"password"`
	FullName string  `json:"full_name" db:"full_name"`
}
