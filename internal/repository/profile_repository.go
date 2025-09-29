package repository

import (
	"context"
	"database/sql"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/siroj05/portfolio/config"
	"github.com/siroj05/portfolio/internal/dto"
)

type ProfileRepsitory struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepsitory {
	return &ProfileRepsitory{db: db}
}

func (r *ProfileRepsitory) GetById(ctx context.Context, res *dto.ResponseProfileDto, id int64) error {
	row := r.db.QueryRowContext(ctx, `SELECT * FROM profile WHERE user_id = ?`, id)
	err := row.Scan(&res.ID, &res.UserID, &res.ImagePath, &res.FullName, &res.JobTitle, &res.Email, &res.Linkedin, &res.Repository, &res.About, &res.PhoneNumber, &res.Location)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	res.ImagePath = config.BaseUrlImg + res.ImagePath

	return nil
}

func (r *ProfileRepsitory) IsFileExist(id int64) (string, error) {
	var oldFilePath string
	err := r.db.QueryRow("SELECT image_path FROM profile WHERE user_id = ?", id).Scan(&oldFilePath)
	if err != nil {
		return "", err
	}
	return oldFilePath, nil
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
			about = ?,
			phone_number = ?,
			location = ?
			WHERE user_id = ?`, req.ImagePath, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About, req.PhoneNumber, req.Location, req.UserID)
		} else {
			_, err = tx.Exec(`
				UPDATE profile
				SET 
				full_name = ?,
				job_title = ?,
				email = ?,
				linkedin = ?,
				repo = ?,
				about = ?,
				phone_number = ?,
				location = ?
				WHERE user_id = ?`, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About, req.PhoneNumber, req.Location, req.UserID)
		}
	} else {
		id := uuid.New().String()
		_, err = tx.Exec(`
		INSERT INTO profile 
		(id, user_id, image_path, full_name, job_title, email, linkedin, repo, about, phone_number, location) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, id, req.UserID, req.ImagePath, req.FullName, req.JobTitle, req.Email, req.Linkedin, req.Repository, req.About, req.PhoneNumber, req.Location)
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

func (r *ProfileRepsitory) Get(ctx context.Context) ([]dto.ResponseProfileDto, error) {

	rows, err := r.db.QueryContext(ctx, `SELECT * FROM profile ORDER BY full_name ASC`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res = make(map[string]*dto.ResponseProfileDto)
	BASE_URL := config.BaseUrlImg
	for rows.Next() {
		var (
			ID          string
			UserID      int64
			ImagePath   string
			FullName    string
			JobTitle    string
			Email       string
			Linkedin    string
			Repository  string
			About       string
			PhoneNumber string
			Location    string
		)
		rows.Scan(&ID, &UserID, &ImagePath, &FullName, &JobTitle, &Email, &Linkedin, &Repository, &About, &PhoneNumber, &Location)
		if _, ok := res[ID]; !ok {
			res[ID] = &dto.ResponseProfileDto{
				ID:          ID,
				UserID:      UserID,
				ImagePath:   BASE_URL + ImagePath,
				FullName:    FullName,
				JobTitle:    JobTitle,
				Email:       Email,
				Linkedin:    Linkedin,
				Repository:  Repository,
				About:       About,
				PhoneNumber: PhoneNumber,
				Location:    Location,
			}
		}
	}

	result := make([]dto.ResponseProfileDto, 0, len(res))
	for _, cat := range res {
		result = append(result, *cat)
	}

	return result, nil
}
