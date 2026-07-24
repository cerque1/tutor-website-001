package model

type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	PasswordHash string `json:"password_hash"`
	IsAdmin bool `json:"is_admin"`
}
