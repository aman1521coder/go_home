package middleware

import (
	"net/http"
	"primeauction/api/utils"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := parts[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		r.Header.Set("X-User-ID", claims.UserID)
		r.Header.Set("X-User-Email", claims.Email)
		if claims.IsAdmin {
			r.Header.Set("X-Is-Admin", "true")
		}
		next(w, r)
	}
}
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminHeader := r.Header.Get("X-Is-Admin")
		if adminHeader != "true" {
			http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
