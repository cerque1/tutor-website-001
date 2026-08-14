package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
	"github.com/cerque1/tutor-website-001/internal/middleware"
)

func registerReviewRouter(mux *http.ServeMux, h *handler.ReviewHandler, secret string) {
	mux.Handle(
		"GET /review/all",
		middleware.Chain(
			http.HandlerFunc(h.GetAll),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"POST /review/",
		middleware.Chain(
			http.HandlerFunc(h.Create),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"GET /review/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Get),
			middleware.Auth(secret),
		),
	)

	mux.Handle(
		"DELETE /review/{id}",
		middleware.Chain(
			http.HandlerFunc(h.Delete),
			middleware.Auth(secret),
		),
	)
}
