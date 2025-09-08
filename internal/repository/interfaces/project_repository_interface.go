package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProjectRepository interface {
	Create(ctx context.Context, project dto.ProjectDto) error
	GetAll(ctx context.Context) ([]dto.ProjectDto, error)
}
