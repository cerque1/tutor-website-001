package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/service"
	"github.com/go-playground/validator/v10"
)

type ServiceHandler struct {
	service *service.ServiceService
	validate *validator.Validate
}

func NewServiceHandler(
	s *service.ServiceService,
	v *validator.Validate,
) *ServiceHandler {
	return &ServiceHandler{
		service: s,
		validate: v,
	}
}

// GetAll godoc
// @Summary Получить список услуг
// @Description Получить весь список услуг
// @Tags Services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Succsess 200 {array} dto.Service
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /services/all [get]
func (h *ServiceHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	services, err := h.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

// Create godoc
// @Summary Создать услугу
// @Description Создать новый сервис
// @Tags Services
// @Accept json
// @Produce json
// @Param request body dto.UserCreate true "Данные услуги"
// @Security BearerAuth
// @Succsess 201 {object} dto.Service
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /services [post]
func (h *ServiceHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.ServiceCreate

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serv, err := h.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(serv); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}

// Get godoc
// @Summary Получить сервис
// @Description Получает сервис по id
// @Tags Services
// @Produce json
// @Param id path uint64 true "ID сервиса"
// @Security BearerAuth
// @Succsess 200 {object} dto.Service
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /services/{id} [get]
func (h *ServiceHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {

	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	serv, err := h.service.Get(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(serv); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
		return
	}
}

// Patch godoc
// @Summary Изменить сервис
// @Description Частичное изменение данных сервиса
// @Tags Services
// @Accept json
// @Produce json
// @Param id path uint64 true "ID услуги"
// @Param request body dto.ServicePatch true "Данные для изменения"
// @Security BearerAuth
// @Succsess 200 {object} dto.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /services/{id} [patch]
func (h *ServiceHandler) Patch(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.ServicePatch

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	serv, err := h.service.Patch(r.Context(), idx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(serv); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
		return
	}
}

// Delete godoc
// @Summary Удалить услугу
// @Tags Services
// @Param id path uint64 true "ID услуги"
// @Security BearerAuth
// @Succsess 204
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /services/{id} [delete]
func (h *ServiceHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(r.Context(), idx)
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
