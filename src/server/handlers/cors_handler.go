package handlers

import (
	"fmt"
	"net/http"

	"github.com/shota-imoto/helixf/lib/utils/domain"
)

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CorsHandler called.")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", domain.FrontendUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	w.WriteHeader(200)
	return
}
