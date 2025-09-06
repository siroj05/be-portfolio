package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
)

type ProjectsHandler struct {
	Repo interfaces.ProjectRepository
}

func NewProjectHandler(repo interfaces.ProjectRepository) *ProjectsHandler {
	return &ProjectsHandler{Repo: repo}
}

func (h *ProjectsHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20) // limit upload size to 10mb
	if err != nil {
		response.Error(w, http.StatusBadRequest, "File too big", err.Error())
		return
	}

	// ambil file
	file, handler, err := r.FormFile("picture")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Error retrieving the file", err.Error())
		return
	}
	defer file.Close()

	// simpan file di folder uploads lokal
	filePath := fmt.Sprintf("./uploads/%s", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Unable to save file", err.Error())
		return
	}

	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error saving file", err.Error())
		return
	}

	req := dto.ProjectDto{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		TechStack:   r.FormValue("techStack"),
		DemoUrl:     r.FormValue("demoUrl"),
		GithubUrl:   r.FormValue("githubUrl"),
		FilePath:    filePath,
	}

	err = h.Repo.Create(ctx, req)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create project", err.Error())
		return
	}

	response.Success(w, "Create project success", req)
}
