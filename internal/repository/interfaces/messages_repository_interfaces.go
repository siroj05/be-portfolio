package interfaces

import (
	"context"

	"github.com/siroj05/portfolio/internal/dto"
)

type MessagesRepository interface {
	Create(ctx context.Context, message dto.MessageDto) error
	// GetAll(ctx context.Context) ([]dto.MessageDto, error)
	// Delete(ctx context.Context, id int) error
}
