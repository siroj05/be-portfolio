package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
)

type ExperiencesHandler struct {
	Repo interfaces.ExperiencesRepository
}

func NewExperiencesHandler(repo interfaces.ExperiencesRepository) *ExperiencesHandler {
	return &ExperiencesHandler{
		Repo: repo,
	}
}

func (h *ExperiencesHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	var req dto.ExperiencesDto
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// format time
	// validasi
	if strings.TrimSpace(req.Office) == "" {
		response.Error(w, http.StatusBadRequest, "Office is required", err.Error())
		return
	}

	if strings.TrimSpace(req.Position) == "" {
		response.Error(w, http.StatusBadRequest, "Position is required", err.Error())
		return
	}

	if strings.TrimSpace(req.Description) == "" {
		response.Error(w, http.StatusBadRequest, "Description is required", err.Error())
		return
	}

	if strings.TrimSpace(req.Start) == "" {
		response.Error(w, http.StatusBadRequest, "Start is required", err.Error())
		return
	}

	// if req.End.IsZero() {
	// 	response.Error(w, http.StatusBadRequest, "End is required", err.Error())
	// 	return
	// }

	err = h.Repo.Create(ctx, req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to create experience", err.Error())
		return
	}

	response.Success(w, "Create experience success", nil)
}

func (h *ExperiencesHandler) GetAllExperiences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	res, err := h.Repo.GetAll(ctx)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to get experiences", err.Error())
		return
	}

	response.Success(w, "Successfully get all experiences", res)
}

func (h *ExperiencesHandler) GetExperienceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]

	var res dto.ExperiencesListDto
	err := h.Repo.GetById(ctx, id, &res)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}

	response.Success(w, "Successfully get data", res)
}

func (h *ExperiencesHandler) DeleteExperiences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]

	err := h.Repo.Delete(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete experience", err.Error())
		return
	}

	response.Success(w, "Successfully deleted message", nil)
}

func (h *ExperiencesHandler) UpadateExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	var req dto.ExperiencesDto
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.Repo.Update(ctx, req)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to update experience", err.Error())
		return
	}

	response.Success(w, "Successfully to update experience", nil)
}
