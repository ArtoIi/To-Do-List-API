package interfaces

import (
	"context"
	"net/http"
	"strings"

	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/security"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Autorização necessária", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
			return
		}

		claims, err := security.ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user_id", claims["id"]))

		next.ServeHTTP(w, r)
	})
}
