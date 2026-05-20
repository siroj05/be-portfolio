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

// LoginUser godoc
// @Summary User Login
// @Description Authenticates user and sets session cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginDto true "User credentials"
// @Success 200 {object} map[string]string "Login success message"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 401 {object} response.Response "Invalid credentials"
// @Router /auth/login [post]
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

// LogoutUser godoc
// @Summary User Logout
// @Description Clears user session cookie
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string "Logout success message"
// @Router /auth/logout [post]
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

// CreateUser godoc
// @Summary Register a User
// @Description Creates a new admin/user
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginDto true "User credentials"
// @Success 200 {object} map[string]string "User registration success message"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 500 {object} response.Response "Failed to create user"
// @Router /auth/register [post]
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

// GetDataUser godoc
// @Summary Get Current Profile
// @Description Retrieves the authenticated user profile details
// @Tags Auth
// @Produce json
// @Success 200 {object} response.Response{data=dto.GetMeDto} "Successfully get profile data"
// @Failure 400 {object} response.Response "Invalid token"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/me [get]
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

// ForgotPassword godoc
// @Summary Forgot Password
// @Description Resets/changes password for a specified user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.ForgotPasswordDto true "Reset password payload"
// @Success 200 {object} map[string]string "Password updated successfully"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 500 {object} response.Response "Failed to reset password"
// @Router /auth/forgot [post]
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req dto.ForgotPasswordDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if req.Name == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "Username and new password are required", nil)
		return
	}

	ctx := context.Background()
	err := h.Repo.ResetPassword(ctx, req)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to reset password", err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password updated successfully",
	})
}

