package controllers

import (
	"encoding/json"
	"goauth/internal/middlewares"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username := claims["username"]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Profile loaded successfully",
		"username": username,
	})
}
