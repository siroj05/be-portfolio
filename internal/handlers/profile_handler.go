package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
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
		log.Println("error 1 = ", err)
		response.Error(w, http.StatusBadRequest, "Failed to parse form", err.Error())
		return
	}

	file, handler, err := r.FormFile("image")
	// convert string ke int
	i64, err := strconv.ParseInt(r.FormValue("userId"), 10, 64)
	if err != nil {
		log.Println("error 5 = ", err)
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	if err != nil && err != http.ErrMissingFile {
		log.Println("error 2 = ", err)
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

		oldFilePath, err := h.Repo.IsFileExist(i64)
		log.Println("error 2.5 = ", err)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		}
		log.Println(oldFilePath)
		if oldFilePath != "" {
			_ = os.Remove(oldFilePath)
		}

		dst, err := os.Create(filePath)
		if err != nil {
			log.Println("error 3 = ", err)
			response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Println("error 4 = ", err)
			response.Error(w, http.StatusInternalServerError, "Failed to write file", err.Error())
			return
		}
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
		log.Println("error 6 = ", err)
		if filePath != "" {
			os.Remove(filePath)
		}
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	response.Success(w, "Successfully to save profile", nil)
}
