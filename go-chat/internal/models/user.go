package models

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Register struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	FullName string `json:"full_name" form:"full_name"`
}

type User struct {
	Avatar   *string `json:"avatar" db:"avatar"`
	ID       string  `json:"id" db:"id"`
	Username string  `json:"username" db:"username"`
	Password string  `json:"password" db:"password"`
	FullName string  `json:"full_name" db:"full_name"`
}
