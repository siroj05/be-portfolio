package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siroj05/portfolio/config"
	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/models"
	"github.com/siroj05/portfolio/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Login(ctx context.Context, req dto.LoginDto) (string, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, password FROM user WHERE name = ?", req.Name)

	var user models.Auth
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("Invalid username")
		}
		return "", err
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("Invalid username or password")
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString([]byte(config.JWTSecret))

	if err != nil {
		return "", nil
	}

	return TokenString, nil
}

func (r *AuthRepository) Create(ctx context.Context, req dto.LoginDto) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, "INSERT INTO user (name, password) VALUES (?, ?)", req.Name, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}
