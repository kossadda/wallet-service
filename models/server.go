package models

import (
	"context"
	"net/http"
	"time"
)

// Custom type with server description
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new server.
// Returns: *Server
func NewServer() *Server {
	return &Server{}
}

// Run starts the server on the given port with the given handler.
// Returns: error
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
// Returns: error
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
