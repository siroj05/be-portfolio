package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProfileRepository interface {
	GetById(ctx context.Context, res *dto.ResponseProfileDto, id int64) error
	Create(ctx context.Context, req dto.ProfileDto) error
	IsFileExist(id int64) (string, error)
	Get(ctx context.Context) ([]dto.ResponseProfileDto, error)
}
