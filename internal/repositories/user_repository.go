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

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT id, username, password_hash, role, created_at FROM users WHERE username = $1",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
