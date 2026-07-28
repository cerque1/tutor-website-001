package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

func registerUserRouters(mux *http.ServeMux, h *handler.UserHandler) {
	mux.HandleFunc(
		"GET /users/all",
		h.GetAll,
	)

	mux.HandleFunc(
		"POST /users/",
		h.Create,
	)

	mux.HandleFunc(
		"GET /users/{id}",
		h.Get,
	)

	mux.HandleFunc(
		"PATCH /users/{id}",
		h.Patch,
	)

	mux.HandleFunc(
		"DELETE /users/{id}",
		h.Delete,
	)
}
