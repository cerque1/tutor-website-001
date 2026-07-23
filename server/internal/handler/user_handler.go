package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (u *UserHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	users, err := u.service.GetAll(r.Context())

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}