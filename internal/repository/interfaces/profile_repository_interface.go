package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProfileRepository interface {
	GetById(ctx context.Context, res dto.ProfileDto, id int64) error
	Create(ctx context.Context, req dto.ProfileDto) error
}
