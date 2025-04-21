package services

import (
	"errors"
	"fmt"
	"goauth/internal/models"
	"goauth/internal/repositories"
	"goauth/internal/utils"
	"time"
)

type AuthService struct {
	UserRepo         *repositories.UserRepository
	RefreshTokenRepo *repositories.RefreshTokenRepository
}

func NewAuthService(userRepo *repositories.UserRepository, refreshTokenRepo *repositories.RefreshTokenRepository) *AuthService {
	return &AuthService{
		UserRepo:         userRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
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

func (s *AuthService) Authenticate(username, password string) (accessToken, refreshToken string, err error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		fmt.Println("❌ Usuário não encontrado")
		return "", "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		fmt.Println("❌ Senha incorreta")
		return "", "", errors.New("invalid username or password")
	}

	accessToken, err = utils.GenerateJWT(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = utils.GenerateRandomToken()
	if err != nil {
		return "", "", err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(7 * 24 * time.Hour)

	err = s.RefreshTokenRepo.Save(user.ID, refreshToken, issuedAt, expiresAt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(refreshToken string) (newAccessToken string, newRefreshToken string, err error) {
	userID, err := s.RefreshTokenRepo.Validate(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	newAccessToken, err = utils.GenerateJWT(user)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err = utils.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	err = s.RefreshTokenRepo.Revoke(refreshToken)
	if err != nil {
		return "", "", err
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(7 * 24 * time.Hour)

	err = s.RefreshTokenRepo.Save(user.ID, newRefreshToken, issuedAt, expiresAt)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}
