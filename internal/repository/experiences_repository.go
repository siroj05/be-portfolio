package repository

import (
	"context"
	"database/sql"
	"log"

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
	var end interface{}
	if req.End.Valid {
		end = req.End.Time.Format("2006-01-02")
	} else {
		end = nil
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO experiences (id, office, position, start, end, description) VALUES (?, ?, ?, ?, ?, ?)", id, req.Office, req.Position, req.Start.Time.Format("2006-01-02"), end, req.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) GetAll(ctx context.Context) ([]dto.ExperiencesListDto, error) {
	log.Println("masuk")
	query := `SELECT id, office, position, description, start FROM experiences`
	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res = make([]dto.ExperiencesListDto, 0)

	for row.Next() {
		var e dto.ExperiencesListDto
		log.Println("masuk")
		row.Scan(
			&e.ID,
			&e.Office,
			&e.Position,
			&e.Description,
			&e.Start,
		)
		res = append(res, e)
	}
	log.Println(res)
	return res, nil
}
