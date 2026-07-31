package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/cerque1/tutor-website-001/internal/security"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	IsAdminKey contextKey = "is_admin"
)

func Auth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(auth, "Bearer ")

			claims, err := security.ParseToken(token, []byte(secret))
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.ID)
			ctx = context.WithValue(r.Context(), IsAdminKey, claims.IsAdmin)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
