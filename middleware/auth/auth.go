package middleware

import (
	"net/http"
	"strings"

	"github.com/nomannaq/webshop-api-golang/cmd/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			http.Error(w, "Token is missing", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims.UserID == 0 {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}

}
