package repository

import (
	"context"
	"database/sql"
	"errors"
	"os"

	"github.com/google/uuid"
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

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			if req.ImagePath != "" {
				os.Remove(req.ImagePath)
			}
			panic(p)
		}
	}()

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM profile WHERE user_id = ?)", req.UserID).Scan(&exists)
	if err != nil {
		tx.Rollback()
		if req.ImagePath != "" {
			os.Remove(req.ImagePath)
		}
		return err
	}

	if exists {
		if req.ImagePath != "" {
			_, err = tx.Exec(`
			UPDATE profile 
			SET 
			image_path = ?, 
			full_name = ?, 
			job_title = ?, 
			email = ?, 
			linkedin = ?, 
			repo = ?, 
			about = ? 
			WHERE user_id = ?`, req.ImagePath, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About, req.UserID)
		} else {
			_, err = tx.Exec(`
				UPDATE profile
				SET 
				full_name = ?,
				job_title = ?,
				email = ?,
				linkedin = ?,
				repo = ?,
				about = ?
				WHERE user_id = ?`, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About, req.UserID)
		}
	} else {
		id := uuid.New().String()
		_, err = tx.Exec(`
		INSERT INTO profile 
		(id, user_id, image_path, full_name, job_title, email, linkedin, repo, about) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, id, req.UserID, req.ImagePath, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About)
	}

	if err != nil {
		tx.Rollback()
		if req.ImagePath != "" {
			os.Remove(req.ImagePath)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if req.ImagePath != "" {
			os.Remove(req.ImagePath)
		}
		return err
	}

	return nil
}
