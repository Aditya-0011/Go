package middlewares

import (
	"net/http"
	"strconv"
	"strings"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"msg":"missing Authorization header"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"msg":"invalid auth format"}`, http.StatusUnauthorized)
			return
		}

		num, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, `{"msg":"invalid token value"}`, http.StatusUnauthorized)
			return
		}

		if num%2 != 0 {
			http.Error(w, `{"msg":"not allowed"}`, http.StatusUnauthorized)
			return
		}
		if isPowerOfTwo(num) {
			next(w, r)
			return
		}
		http.Error(w, `{"msg":"token expired"}`, http.StatusForbidden)
	}
}