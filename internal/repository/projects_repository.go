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
	(title, description, tech_stack, demo_url, github_url, filepath) 
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, req.Title, req.Description, req.TechStack, req.DemoUrl, req.GithubUrl, req.FilePath)

	if err != nil {
		return err
	}

	return nil
}
