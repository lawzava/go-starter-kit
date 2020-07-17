package middleware

import (
	"goapp/pkg/logger"
	"net/http"
)

func Logger(log *logger.Log, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.
			With("method", r.Method).
			Info(r.URL.String())

		next.ServeHTTP(w, r)
	})
}
