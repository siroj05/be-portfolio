package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/siroj05/portfolio/internal/dto"
)

type SkillsRepository struct {
	db *sql.DB
}

func NewSkillsRepository(db *sql.DB) *SkillsRepository {
	return &SkillsRepository{db: db}
}

func (r *SkillsRepository) Create(ctx context.Context, req dto.CategoriesDto) error {
	CategoryID := uuid.New().String()
	_, err := r.db.ExecContext(ctx, "INSERT INTO categories (id, name) VALUES (?, ?)", CategoryID, req.Category)
	if err != nil {
		return err
	}

	for _, skill := range req.Skills {
		skillID := uuid.New().String()
		_, err := r.db.ExecContext(ctx, "INSERT INTO SKILLS (id, category_id, name, icon) VALUES (?, ?, ?, ?)", skillID, CategoryID, skill.Name, skill.Icon)
		if err != nil {
			return err
		}
	}

	return nil

}
