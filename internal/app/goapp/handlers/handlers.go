package handlers

import (
	"goapp/pkg/logger"
	"net/http"
)

type Handlers interface {
	Handle(pattern string, handler http.Handler)
}

func Register(h Handlers, log *logger.Log) {
	h.Handle("/ping", handlePing(log))
}
