package repository

import (
	"context"
	"database/sql"

	"github.com/siroj05/portfolio/internal/dto"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, req dto.ProjectDto) error {
	query := `
	INSERT INTO projects 
	(id, title, description, tech_stack, demo_url, github_url, filepath) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, req.ID, req.Title, req.Description, req.TechStack, req.DemoUrl, req.GithubUrl, req.FilePath)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectRepository) GetAll(ctx context.Context) ([]dto.ProjectDto, error) {
	query := `SELECT * FROM projects`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	BASE_URL := "http://localhost:8080/"
	var res = make([]dto.ProjectDto, 0)
	for rows.Next() {
		var r dto.ProjectDto
		rows.Scan(&r.ID, &r.Title, &r.Description, &r.TechStack, &r.DemoUrl, &r.GithubUrl, &r.FilePath)
		res = append(res,
			dto.ProjectDto{
				ID:          r.ID,
				Title:       r.Title,
				Description: r.Description,
				TechStack:   r.TechStack,
				DemoUrl:     r.DemoUrl,
				GithubUrl:   r.GithubUrl,
				FilePath:    BASE_URL + r.FilePath,
			},
		)
	}

	return res, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM projects WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return err
}
