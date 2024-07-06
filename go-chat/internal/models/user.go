package models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	FullName string  `json:"full_name"`
	Avartar  *string `json:"avatar"`
}
