package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type ExperiencesRepository interface {
	Create(ctx context.Context, req dto.ExperiencesDto) error
}
