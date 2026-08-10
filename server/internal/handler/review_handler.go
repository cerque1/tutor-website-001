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
