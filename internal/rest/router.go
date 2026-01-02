package rest

import (
	"encoding/json"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/users", handleUsers)
	return JSONMiddleware(mux)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"message": "Welcome to the Go REST API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"users": "/users",
		},
	}

	json.NewEncoder(w).Encode(response)
}
