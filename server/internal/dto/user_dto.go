package dto

type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	IsAdmin bool `json:"is_admin"`
}

type UserCreate struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type UserPatch struct {
	Name *string `json:"name" validate:"omitempty,min=3,max=50"`
	Password *string `json:"password" validate:"omitempty,min=8,max=72"`
}
