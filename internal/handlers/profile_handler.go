package handlers

import (
	"net/http"

	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
)

type ProfileHandler struct {
	Repo interfaces.ProfileRepository
}

func NewProfileHandler(repo interfaces.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{Repo: repo}
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to parse form", err.Error())
		return
	}

	response.Success(w, "Create profile successfully", nil)

}
