package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	ctx := context.Background()

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20) // limit upload size to 10mb
	if err != nil {
		log.Println("error 1")
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "File too big", err.Error())
		return
	}

	// ambil file
	file, handler, err := r.FormFile("image")

	if err != nil {
		log.Println("error 2")
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Error retrieving the file", err.Error())
		return
	}
	defer file.Close()

	// simpan file di folder uploads lokal
	filePath := fmt.Sprintf("uploads/%s", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		log.Println("error 3")
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Unable to save file", err.Error())
		return
	}

	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println("error 4")
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Error saving file", err.Error())
		return
	}
	id := uuid.New().String()
	req := dto.ProjectDto{
		ID:          id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		TechStack:   r.FormValue("techStack"),
		DemoUrl:     r.FormValue("demoUrl"),
		GithubUrl:   r.FormValue("githubUrl"),
		FilePath:    filePath,
	}

	err = h.Repo.Create(ctx, req)

	if err != nil {
		log.Println("error 5")
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to create project", err.Error())
		return
	}

	response.Success(w, "Create project success", req)
}

func (h *ProjectsHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	res, err := h.Repo.GetAll(ctx)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to get projects", err.Error())
		return
	}

	response.Success(w, "Successfully get all projects", res)
}

func (h *ProjectsHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	err := h.Repo.Delete(ctx, id)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to delete project", err.Error())
		return
	}

	response.Success(w, "Successfully deleted project", nil)
}

func (h *ProjectsHandler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	var res dto.ProjectDto
	err := h.Repo.GetById(ctx, id, &res)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	response.Success(w, "Successfully get data", res)
}

func (h *ProjectsHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// ambil id dari URL
	vars := mux.Vars(r)
	id := vars["id"]

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20) // buat ngelimit size
	if err != nil {
		log.Println("error 1")
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "File too big", err.Error())
		return
	}

	// ambil file
	req := dto.ProjectDto{
		ID:          id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		TechStack:   r.FormValue("techStack"),
		DemoUrl:     r.FormValue("demoUrl"),
		GithubUrl:   r.FormValue("githubUrl"),
		FilePath:    "",
	}

	// coba ambil file (optional)
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// simpan file baru
		filePath := fmt.Sprintf("uploads/%s", handler.Filename)
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

		req.FilePath = filePath // set filepath kalau ada file baru
	}

	// update ke repo
	err = h.Repo.Update(ctx, req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update project", err.Error())
		return
	}

	response.Success(w, "Update project success", req)
}
