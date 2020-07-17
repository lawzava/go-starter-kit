package handlers

import (
	"goapp/pkg/logger"
	"net/http"
)

func handlePing(log *logger.Log) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("logging ping handler")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	}
}
