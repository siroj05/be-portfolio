package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type SkillsRepository interface {
	Create(ctx context.Context, req dto.CategoriesDto) error
}
