package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/siroj05/portfolio/internal/dto"
)

type ExperiencesRepository struct {
	db *sql.DB
}

func NewExperiencesRepository(db *sql.DB) *ExperiencesRepository {
	return &ExperiencesRepository{
		db: db,
	}
}

func (r *ExperiencesRepository) Create(ctx context.Context, req dto.ExperiencesDto) error {
	id := uuid.New().String()
	_, err := r.db.ExecContext(ctx, "INSERT INTO experiences (id, office, position, start, end, description) VALUES (?, ?, ?, ?, ?, ?)", id, req.Office, req.Position, req.Start, req.End, req.Description)
	if err != nil {
		return err
	}

	return nil
}
