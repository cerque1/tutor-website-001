package dto

type Review struct {
	ID      uint64 `json:"id"`
	UserID  uint64 `json:"user_id"`
	Content string `json:"content"`
	Rating  uint16 `json:"rating"`
}

type ReviewCreate struct {
	Content string `json:"content" validate:"required,min=10,max=1000"`
	Rating  uint16 `json:"rating" validate:"required,gte=1,lte=5"`
}
