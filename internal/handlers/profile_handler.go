package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/siroj05/portfolio/internal/dto"
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
	ctx := context.Background()
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to parse form", err.Error())
		return
	}

	file, handler, err := r.FormFile("image")

	if err != nil && err != http.ErrMissingFile {
		response.Error(w, http.StatusBadRequest, "Failed to read the file", err.Error())
		return
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	var filePath string
	if file != nil {

		// rename file pakai timestamp
		newFileName := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), handler.Filename)
		filePath = newFileName

		dst, err := os.Create(filePath)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to write file", err.Error())
			return
		}
	}

	// convert string ke int
	i64, err := strconv.ParseInt(r.FormValue("userId"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	req := dto.ProfileDto{
		ID:         r.FormValue("id"),
		UserID:     int64(i64),
		ImagePath:  filePath,
		FullName:   r.FormValue("fullName"),
		JobTitle:   r.FormValue("jobTitle"),
		Email:      r.FormValue("email"),
		Linkedin:   r.FormValue("linkedin"),
		Repository: r.FormValue("repository"),
		About:      r.FormValue("about"),
	}
	err = h.Repo.Create(ctx, req)

	if err != nil {
		if filePath != "" {
			os.Remove(filePath)
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	response.Success(w, "Successfully to save profile", nil)
}
