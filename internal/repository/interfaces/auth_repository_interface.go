package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type AuthRepository interface {
	Create(ctx context.Context, req dto.LoginDto) error
	Login(ctx context.Context, req dto.LoginDto) (string, error)
}
