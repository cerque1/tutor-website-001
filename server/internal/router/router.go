package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

type Handlers struct {
	User *handler.UserHandler
}

func New(h Handlers) http.Handler {
	mux := http.NewServeMux()

	registerUserRouters(mux, h.User)

	return mux
}