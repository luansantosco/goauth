package controllers

import (
	"encoding/json"
	"goauth/internal/services"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, authService *services.AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = authService.Register(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request, s *services.AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := s.Authenticate(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshHandler(w http.ResponseWriter, r *http.Request, s *services.AuthService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type refreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	var req refreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.RefreshToken == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newAccessToken, newRefreshToken, err := s.Refresh(req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
