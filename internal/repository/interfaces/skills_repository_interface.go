package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type SkillsRepository interface {
	Create(ctx context.Context, req dto.CategoriesDto) error
	GetAll(ctx context.Context) ([]dto.CategoriesDto, error)
	Delete(ctx context.Context, id string) error
	GetSkillsByCategory(ctx context.Context, id string) ([]dto.SkillDto, error)
}
