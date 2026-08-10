package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/service"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service *service.UserService
	validate *validator.Validate
}

func NewUserHandler(
	s *service.UserService,
	v *validator.Validate,
) *UserHandler {
	return &UserHandler{
		service: s,
		validate: v,
	}
}

// GetAll godoc
// @Summary Получить список пользователей
// @Description Получить весь список пользователей
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Succsess 200 {array} dto.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/all [get]
func (u *UserHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	users, err := u.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Create godoc
// @Summary Создать пользователя
// @Description Создать нового пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserCreate true "Данные пользователя"
// @Security BearerAuth
// @Succsess 201 {object} dto.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func (u *UserHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.UserCreate

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := u.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}

// Get godoc
// @Summary Получить пользователя
// @Description Получает пользователя по id
// @Tags Users
// @Produce json
// @Param id path uint64 true "ID пользователя"
// @Security BearerAuth
// @Succsess 200 {object} dto.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [get]
func (u *UserHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := u.service.Get(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}

// Patch godoc
// @Summary Изменить пользователя
// @Description Частичное изменение данных пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint64 true "ID пользователя"
// @Param request body dto.UserPatch true "Данные для изменения"
// @Security BearerAuth
// @Succsess 200 {object} dto.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [patch]
func (u *UserHandler) Patch(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err :=	strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.UserPatch

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := u.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.service.Patch(r.Context(), idx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete godoc
// @Summary Удалить пользователя
// @Tags Users
// @Param id path uint64 true "id пользователя"
// @Security BearerAuth
// @Succsess 204 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [delete]
func (u *UserHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = u.service.Delete(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}
