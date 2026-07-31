package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

func registerAuthRouter(mux *http.ServeMux, h *handler.AuthHandler) {
	mux.HandleFunc(
		"GET /auth/",
		h.Login,
	)
}
