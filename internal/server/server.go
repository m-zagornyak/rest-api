package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) RunServer(port string) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpserver.ListenAndServe()
}
