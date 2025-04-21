package repositories

import (
	"database/sql"
	"time"
)

type RefreshTokenRepository struct {
	DB *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{DB: db}
}

func (r *RefreshTokenRepository) Save(userID int, token string, issuedAt, expiresAt time.Time) error {
	_, err := r.DB.Exec(`
		INSERT INTO refresh_tokens (user_id, token, issued_at, expires_at)
		VALUES ($1, $2, $3, $4)
	`, userID, token, issuedAt, expiresAt)
	return err
}

func (r *RefreshTokenRepository) Validate(token string) (int, error) {
	var userID int
	var expiresAt time.Time
	err := r.DB.QueryRow(`
		SELECT user_id, expires_at FROM refresh_tokens WHERE token = $1
	`, token).Scan(&userID, &expiresAt)

	if err != nil {
		return 0, err
	}

	if time.Now().After(expiresAt) {
		return 0, sql.ErrNoRows
	}

	return userID, nil
}

func (r *RefreshTokenRepository) Revoke(token string) error {
	_, err := r.DB.Exec(`
		DELETE FROM refresh_tokens WHERE token = $1
	`, token)
	return err
}
