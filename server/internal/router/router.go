package router

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/cerque1/tutor-website-001/docs"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

type Handlers struct {
	User *handler.UserHandler
	Auth *handler.AuthHandler
	Service *handler.ServiceHandler
	Review *handler.ReviewHandler
}

func New(h Handlers, secret string) http.Handler {
	mux := http.NewServeMux()

	registerUserRouters(mux, h.User, secret)
	registerServiceRouter(mux, h.Service, secret)
	registerAuthRouter(mux, h.Auth)
	registerReviewRouter(mux, h.Review, secret)

	mux.Handle(
		"GET /docs/",
		httpSwagger.WrapHandler,
	)

	return mux
}
