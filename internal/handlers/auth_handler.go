package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/middleware"
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
		response.Error(w, http.StatusUnauthorized, err.Error(), err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
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

func (h *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Logout success"}`))
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	ctx := context.Background()
	err := h.Repo.Create(ctx, req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

func (h *AuthHandler) GetDataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	claims, ok := middleware.GetClaims(r)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized", ok)
		return
	}

	userId, ok := claims["userId"].(float64)
	if !ok {
		response.Error(w, http.StatusBadRequest, "Invalid token", ok)
		return
	}

	ctx := context.Background()

	var res dto.GetMeDto

	err := h.Repo.GetMe(ctx, &res, int64(userId))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	response.Success(w, "Successfully get profile data", res)
}
