package router

import (
	"net/http"

	"github.com/cerque1/tutor-website-001/internal/handler"
)

func NewUserRouter(user *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /users/all",
		user.GetAll,
	)
	return mux
}