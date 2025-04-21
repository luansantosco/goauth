package services

import (
	"goauth/internal/models"
	"goauth/internal/repositories"
	"goauth/internal/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     username,
		PasswordHash: hashedPassword,
		Role:         "user",
	}
	return s.UserRepo.Save(user)
}
