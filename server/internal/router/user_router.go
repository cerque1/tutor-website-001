package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
	"github.com/cerque1/tutor-website-001/internal/middleware"
)

func registerUserRouters(mux *http.ServeMux, h *handler.UserHandler, secret string) {
	mux.Handle(
		"GET /users/all",
		middleware.Chain(
			http.HandlerFunc(h.GetAll),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"POST /users/",
		middleware.Chain(
			http.HandlerFunc(h.Create),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"GET /users/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Get),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"PATCH /users/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Patch),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"DELETE /users/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Delete),
			middleware.Auth(secret),
		),
	)
}
