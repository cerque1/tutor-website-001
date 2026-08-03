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

func (s *ServiceHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	services, err := s.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func (s *ServiceHandler) Create(
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
	if err := s.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serv, err := s.service.Create(r.Context(), req)
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

func (s *ServiceHandler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {

	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	serv, err := s.service.Get(r.Context(), idx)
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

func (s *ServiceHandler) Patch(
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

	serv, err := s.service.Patch(r.Context(), idx, req)
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

func (s *ServiceHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	idx, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = s.service.Delete(r.Context(), idx)
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
