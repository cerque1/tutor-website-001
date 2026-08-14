package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/service"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service *service.AuthService
	validate *validator.Validate
}

func NewAuthHandler(
	s *service.AuthService,
	v *validator.Validate,
) *AuthHandler {
	return &AuthHandler{
		service: s,
		validate: v,
	}
}

// Login godoc
// @Summary Авторизация полльзователя
// @Tags Auth
// @Param request body dto.Login true "Данные аутентификации"
// @Security BearerAuth
// @Succsess 200 {object} dto.TokenResponse
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /review/{id} [delete]
func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.Login

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenStr, err := h.service.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := dto.TokenResponse{Token: tokenStr}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}
