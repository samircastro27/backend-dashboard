package middlewares

import (
	"net/http"
	"os"

	cphttp "github.com/samircastro27/backend-dashboard/cmd/api/http"
)

func APIKeyMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("x-api-key")
			if apiKey == "" {
				cphttp.JSONResponse(w, http.StatusUnauthorized, &cphttp.APIResponse{
					Success: false,
					Data:    nil,
					Error:   "API key not provided",
				})
				return
			}

			if apiKey != os.Getenv("API_KEY") {
				cphttp.JSONResponse(w, http.StatusUnauthorized, &cphttp.APIResponse{
					Success: false,
					Data:    nil,
					Error:   "Invalid API key",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
