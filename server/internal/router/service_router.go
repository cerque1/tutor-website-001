package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
	"github.com/cerque1/tutor-website-001/internal/middleware"
)

func registerServiceRouter(mux *http.ServeMux, h *handler.ServiceHandler, secret string) {
	mux.Handle(
		"GET /service/all",
		middleware.Chain(
			http.HandlerFunc(h.GetAll),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"POST /service/",
		middleware.Chain(
			http.HandlerFunc(h.Create),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"GET /service/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Get),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"PATCH /service/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Patch),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"DELETE /service/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Delete),
			middleware.Auth(secret),
		),
	)
}
