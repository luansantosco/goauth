package controllers

import (
	"encoding/json"
	"goauth/internal/services"
	"net/http"
)

type RegisterResquest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, authService *services.AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "🛑 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterResquest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, "🛑 Invalid request data", http.StatusBadRequest)
		return
	}

	err = authService.Register(req.Username, req.Password)

	if err != nil {
		http.Error(w, "🛑 Failed to register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "✅ User registered successfully"})

}
