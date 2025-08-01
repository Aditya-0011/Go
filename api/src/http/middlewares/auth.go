package middlewares

import (
	"api/src/utils"
	"net/http"
	"strconv"
	"strings"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ErrorResponse{Msg: "Missing Authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ErrorResponse{Msg: "Invalid auth format"})
			return
		}

		num, err := strconv.Atoi(parts[1])
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ErrorResponse{Msg: "Invalid token value"})
			return
		}

		if num%2 != 0 {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ErrorResponse{Msg: "Not allowed"})
			return
		}
		if isPowerOfTwo(num) {
			next.ServeHTTP(w, r)
			return
		}
		utils.WriteJSON(w, http.StatusForbidden, utils.ErrorResponse{Msg: "Token expired"})
	})
}
