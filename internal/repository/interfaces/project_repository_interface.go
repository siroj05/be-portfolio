package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProjectRepository interface {
	Create(ctx context.Context, project dto.ProjectDto) error
	GetAll(ctx context.Context) ([]dto.ProjectDto, error)
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string, res *dto.ProjectDto) error
	Update(ctx context.Context, req dto.ProjectDto) error
}
