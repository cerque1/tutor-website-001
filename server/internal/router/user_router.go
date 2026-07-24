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
}
