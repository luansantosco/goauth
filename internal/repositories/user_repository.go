package repositories

import (
	"database/sql"
	"goauth/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRespository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Save(user *models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)",
		user.Username,
		user.PasswordHash,
		user.Role,
	)
	return err
}
