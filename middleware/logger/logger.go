package loggingmiddleware

import (
	"film-service/proxy/logger"
	"net/http"
)

//Middleware --> Logs the request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("HTTP %s %v", r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
