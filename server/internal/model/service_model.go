package model

type Service struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}
