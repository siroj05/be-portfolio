package repository

import (
	"context"
	"database/sql"
	"strings"

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

func (r *ExperiencesRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM experiences WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) Create(ctx context.Context, req dto.ExperiencesDto) error {
	id := uuid.New().String()

	var end interface{}
	if strings.TrimSpace(req.End) == "" {
		end = nil
	} else {
		end = req.End
	}

	_, err := r.db.ExecContext(ctx, "INSERT INTO experiences (id, office, position, start, end, description) VALUES (?, ?, ?, ?, ?, ?)", id, req.Office, req.Position, req.Start, end, req.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) Update(ctx context.Context, req dto.ExperiencesDto) error {
	query := `
	UPDATE experiences 
	SET 
	office = ?,
	position = ?,
	start = ?,
	end = ?,
	description = ?
	WHERE id = ?
	`
	var end interface{}
	if strings.TrimSpace(req.End) == "" {
		end = nil
	} else {
		end = req.End
	}
	_, err := r.db.ExecContext(ctx, query, req.Office, req.Position, req.Start, end, req.Description, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) GetById(ctx context.Context, id string, req *dto.ExperiencesListDto) error {
	var end sql.NullString
	row := r.db.QueryRowContext(ctx, "SELECT * FROM experiences WHERE id = ?", id)
	err := row.Scan(&req.ID, &req.Office, &req.Position, &req.Start, &end, &req.Description)
	if err != nil {
		return err
	}

	if end.Valid {
		req.End = end.String
	} else {
		req.End = ""
	}

	return nil
}

func (r *ExperiencesRepository) GetAll(ctx context.Context) ([]dto.ExperiencesListDto, error) {
	query := `SELECT id, office, position, description, start, end FROM experiences ORDER BY start DESC`
	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res = make([]dto.ExperiencesListDto, 0)

	for row.Next() {
		var e dto.ExperiencesListDto
		row.Scan(
			&e.ID,
			&e.Office,
			&e.Position,
			&e.Description,
			&e.Start,
			&e.End,
		)
		res = append(res, e)
	}
	return res, nil
}
