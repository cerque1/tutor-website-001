package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cerque1/tutor-website-001/internal/dto"
	"github.com/cerque1/tutor-website-001/internal/service"
	"github.com/go-playground/validator/v10"
)

type ReviewHandler struct {
	service  *service.ReviewService
	validate *validator.Validate
}

func NewReviewHandler(
	s *service.ReviewService,
	v *validator.Validate,
) *ReviewHandler {
	return &ReviewHandler{
		service: s,
		validate: v,
	}
}

// GetAll godoc
// @Summary Получить список отзывов
// @Description Получить весть список отзывов
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Succsess 200 {array} dto.Review
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /review/all [get]
func (h *ReviewHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query()

	limit := 20
	offset := 0

	if value := query.Get("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		limit = parsed
	}

	if value := query.Get("offset"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}
		offset = parsed
	}

	reviews, err := h.service.GetAll(r.Context(), uint64(limit), uint64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

// Create godoc
// @Summary Создать отзыв
// @Description Создать новый отзыв
// @Tags Reviews
// @Accept json
// @Produce json
// @Param request body dto.UserCreate true "Данные отзыва"
// @Security BearerAuth
// @Succsess 201 {object} dto.Review
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /review [post]
func (h *ReviewHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req dto.ReviewCreate

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	review, err := h.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(review); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}

// Get godoc
// @Summary Получить отзыв
// @Description Получиет отзыв по id
// @Tags Reviews
// @Produce json
// @Param id path uint64 true "ID отзыва"
// @Security BearerAuth
// @Succsess 200 {object} dto.Review
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /review/{id} [get]
func (h *ReviewHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	review, err := h.service.Get(r.Context(), idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(review); err != nil {
		http.Error(w, "failed to encode responce", http.StatusInternalServerError)
	}
}

// Delete godoc
// @Summary Удалить отзыв
// @Tags Reviews
// @Param id path uint64 true "id отзыва"
// @Security BearerAuth
// @Succsess 204 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// Router /review [delete]
func (h *ReviewHandler) Delete(
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
