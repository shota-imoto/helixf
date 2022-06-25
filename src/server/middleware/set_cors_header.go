package middleware

import (
	"net/http"

	"github.com/shota-imoto/helixf/lib/utils/domain"
)

func SetCorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		next.ServeHTTP(w, r)
	})
}
