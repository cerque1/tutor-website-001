package model

type Review struct {
	ID uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	Content string `json:"content"`
	Rating uint16 `json:"rating"`
}
