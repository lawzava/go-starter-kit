package server

import (
	"goapp/internal/pkg/middleware"
	"goapp/pkg/logger"
	"net/http"
)

type Server interface {
	Handle(pattern string, handler http.Handler)
	Run(address string) error
}

type server struct {
	mux *http.ServeMux
	log *logger.Log
}

func NewServer(log *logger.Log) Server {
	return &server{
		mux: http.NewServeMux(),
		log: log,
	}
}

func (s *server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, middleware.Logger(s.log, handler))
}

func (s *server) Run(address string) error {
	s.log.Infof("server starting at %s", address)
	return http.ListenAndServe(address, s.mux)
}
