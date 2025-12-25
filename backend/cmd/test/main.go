package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// API test endpoints
	r.HandleFunc("/api/v1/items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Items API endpoint",
			"status":  "working",
			"time":    time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	r.HandleFunc("/api/v1/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Categories API endpoint",
			"status":  "working",
			"time":    time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	r.HandleFunc("/api/v1/statistics/overview", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"fresh_count":        0,
			"expiring_soon_count": 0,
			"expired_count":       0,
			"new_items_count":     0,
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Apply CORS middleware
	handler := corsMiddleware(r)

	port := ":8000"
	fmt.Printf("🚀 测试服务器启动在 http://localhost%s\n", port)
	fmt.Println("")
	fmt.Println("可用的测试接口:")
	fmt.Println("  GET  http://localhost:8000/health")
	fmt.Println("  GET  http://localhost:8000/api/v1/items")
	fmt.Println("  GET  http://localhost:8000/api/v1/categories")
	fmt.Println("  GET  http://localhost:8000/api/v1/statistics/overview")
	fmt.Println("")
	fmt.Println("按 Ctrl+C 停止服务器")

	log.Fatal(http.ListenAndServe(port, handler))
}

