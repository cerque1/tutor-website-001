package dto

type Service struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

type ServiceCreate struct {
	Title string `json:"title" validate:"required,min=3,max=128"`
	Description string `json:"description" validate:"required,min=10,max=2048"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

type ServicePatch struct {
	Title       *string  `json:"title" validate:"omitempty,min=3,max=100"`
	Description *string  `json:"description" validate:"omitempty,min=10,max=2000"`
	Price       *float64 `json:"price" validate:"omitempty,gt=0"`
}
