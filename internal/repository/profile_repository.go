package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProfileRepsitory struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepsitory {
	return &ProfileRepsitory{db: db}
}

func (r *ProfileRepsitory) GetById(ctx context.Context, res *dto.ProfileDto, id int64) error {
	row := r.db.QueryRowContext(ctx, `SELECT * FROM profile WHERE user_id = ?`, id)
	err := row.Scan(&res.ID, &res.UserID, &res.ImagePath, &res.FullName, &res.JobTitle, &res.Email, &res.Linkedin, &res.Repository, &res.About)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	BASE_URL := "http://localhost:8080/"
	res.ImagePath = BASE_URL + res.ImagePath

	return nil
}

func (r *ProfileRepsitory) Create(ctx context.Context, req dto.ProfileDto) error {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM profile WHERE user_id = ?)", req.UserID).Scan(&exists)
	if err != nil {
		return err
	}

	// if exists {
	// 	// ini buat update
	// 	_, err = r.db.ExecContext("UPDATE profile")
	// }
	return nil
}
