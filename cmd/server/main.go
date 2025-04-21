package main

import (
	"goauth/configs"
	"goauth/internal/controllers"
	"goauth/internal/middlewares"
	"goauth/internal/repositories"
	"goauth/internal/services"
	"log"
	"net/http"
)

func main() {
	configs.LoadEnv()
	db := configs.ConnectDB()
	defer db.Close()

	userRepository := repositories.NewUserRespository(db)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(db)
	authService := services.NewAuthService(userRepository, refreshTokenRepo)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.RegisterHandler(w, r, authService)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginHandler(w, r, authService)
	})

	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		controllers.RefreshHandler(w, r, authService)
	})

	http.HandleFunc("/logout", controllers.LogoutHandler)

	http.HandleFunc("/profile", middlewares.JWTMiddleware(controllers.ProfileHandler))

	log.Println("✅ Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
