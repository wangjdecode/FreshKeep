package main

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"

	"github.com/freshkeep/backend/internal/conf"
)

// NewHTTPServer creates a new HTTP server with routes
func NewHTTPServer(cfg *conf.Server) *http.Server {
	r := mux.NewRouter()
	
	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	
	// API routes placeholder
	r.PathPrefix("/api/v1/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API endpoint - Protobuf handlers will be registered here"}`))
	})

	return http.NewServer(
		http.Address(cfg.GetHttp().GetAddr()),
		http.Handler(r),
	)
}

