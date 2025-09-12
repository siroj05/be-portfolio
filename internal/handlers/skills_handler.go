package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
)

type SkillsHandler struct {
	Repo interfaces.SkillsRepository
}

func NewSkillsHandler(repo interfaces.SkillsRepository) *SkillsHandler {
	return &SkillsHandler{Repo: repo}
}

func (h *SkillsHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "File too big", err.Error())
		return
	}

	category := r.FormValue("category")

	var uploadedFiles []string
	var skills []dto.SkillDto

	for i := 0; ; i++ {
		name := r.FormValue(fmt.Sprintf("skills[%d][name]", i))
		if name == "" {
			break
		}

		var filePath string
		file, header, err := r.FormFile(fmt.Sprintf("skills[%d][icon]", i))
		if err == nil {
			defer file.Close()
			timestamp := time.Now().Unix()
			fileName := fmt.Sprintf("%d_%s", timestamp, header.Filename)
			filePath = "uploads/" + fileName
			dst, _ := os.Create(filePath)
			defer dst.Close()
			io.Copy(dst, file)
			uploadedFiles = append(uploadedFiles, filePath)
		}

		skills = append(skills, dto.SkillDto{
			Name: name,
			Icon: filePath,
		})
	}

	// simpan
	err = h.Repo.Create(ctx, dto.CategoriesDto{
		Category: category,
		Skills:   skills,
	})

	if err != nil {
		log.Println(err)
		for _, path := range uploadedFiles {
			_ = os.Remove(path)
		}
		response.Error(w, http.StatusBadRequest, "Failed to save", err.Error())
		return
	}

	response.Success(w, "Create project success", nil)
}

func (h *SkillsHandler) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	res, err := h.Repo.GetAll(ctx)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}

	response.Success(w, "Successfully to get all skills", res)
}

func (h *SkillsHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]

	// ente ambil dulu semua skill nya sebelum di hapus di db
	rows, err := h.Repo.GetSkillsByCategory(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}

	// hapus kategori dan skill di db
	err = h.Repo.Delete(ctx, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	// hapus file
	for _, skill := range rows {
		if skill.Icon != "" {
			os.Remove(skill.Icon)
		}
	}

	response.Success(w, "Delete successfully", nil)
}
