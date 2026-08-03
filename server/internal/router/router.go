package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

type Handlers struct {
	User *handler.UserHandler
	Auth *handler.AuthHandler
	Service *handler.ServiceHandler
}

func New(h Handlers, secret string) http.Handler {
	mux := http.NewServeMux()

	registerUserRouters(mux, h.User, secret)
	registerServiceRouter(mux, h.Service, secret)
	registerAuthRouter(mux, h.Auth)

	return mux
}
