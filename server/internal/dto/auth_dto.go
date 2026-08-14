package dto

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type AuthUser struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	IsAdmin bool `json:"is_admin"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
