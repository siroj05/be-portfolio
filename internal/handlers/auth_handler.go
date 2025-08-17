package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
)

type AuthHandler struct {
	Repo interfaces.AuthRepository
}

func NewAuthHandler(repo interfaces.AuthRepository) *AuthHandler {
	return &AuthHandler{
		Repo: repo,
	}
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	ctx := context.Background()
	TokenString, err := h.Repo.Login(ctx, req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusUnauthorized, "Wrong name or password", err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "__session",
		Value:    TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   2 * 60 * 60,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login success",
	})
}
