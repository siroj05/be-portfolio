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
			filePath = "uploads/" + header.Filename
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
		for _, path := range uploadedFiles {
			_ = os.Remove(path)
		}
		response.Error(w, http.StatusBadRequest, "Failed to save", err.Error())
		return
	}

	response.Success(w, "Create project success", nil)
}
