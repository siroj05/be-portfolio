package repository

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/siroj05/portfolio/config"
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
	BASE_URL := config.BaseUrlImg
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

// next handle kalo delete gagal gambar tetap ada
func (r *ProjectRepository) Delete(ctx context.Context, id string) error {

	/* fix gambar ga ke hapus di storage
	*Ambil dulu path nya dari db
	*Terus hapus sesuai path dan name file nya
	 */
	var filePath string
	err := r.db.QueryRowContext(ctx, "SELECT filepath FROM projects WHERE id = ?", id).Scan(&filePath)
	if err != nil {
		return err
	}

	// fungsi hapus file
	_, err = r.db.ExecContext(ctx, `DELETE FROM projects WHERE id = ?`, id)
	if err != nil {
		return err
	}
	if filePath != "" {
		err = os.Remove(filePath)
		if err != nil && !os.IsNotExist(err) {
			// balikin error, note : error bukan karena file ga ada
			return err
		}
	}
	return err
}

func (r *ProjectRepository) GetById(ctx context.Context, id string, res *dto.ProjectDto) error {
	row := r.db.QueryRowContext(ctx, `SELECT * FROM projects WHERE id = ?`, id)
	err := row.Scan(&res.ID, &res.Title, &res.Description, &res.TechStack, &res.DemoUrl, &res.GithubUrl, &res.FilePath)
	if err != nil {
		return err
	}
	BASE_URL := config.BaseUrlImg
	res.FilePath = BASE_URL + res.FilePath

	return nil
}

func (r *ProjectRepository) Update(ctx context.Context, req dto.ProjectDto) error {
	// kalo update ambil dulu data lama
	var oldFilePath string
	err := r.db.QueryRowContext(ctx, `SELECT filepath FROM projects WHERE id = ?`, req.ID).Scan(&oldFilePath)
	if err != nil {
		log.Println("error disini")
		return err
	}
	// ni kalo gambar di ganti
	if req.FilePath != "" {
		// lu hapus ni file lama disini fungsinya
		if oldFilePath != "" {
			_ = os.Remove(oldFilePath)
		}

		// disini updatenya
		_, err = r.db.ExecContext(ctx,
			`UPDATE projects
			SET 
			title = ?, 
			description = ?, 
			tech_stack = ?, 
			demo_url = ?, 
			github_url = ?, 
			filepath = ?
			WHERE id = ?
			`,
			req.Title, req.Description, req.TechStack, req.DemoUrl, req.GithubUrl, req.FilePath, req.ID,
		)
		if err != nil {
			return err
		}
	} else {
		// ni buat kalo lo update tapi ga ganti gambar
		_, err = r.db.ExecContext(ctx,
			`UPDATE projects
			SET 
			title = ?, 
			description = ?, 
			tech_stack = ?, 
			demo_url = ?, 
			github_url = ?
			WHERE id = ?
			`,
			req.Title, req.Description, req.TechStack, req.DemoUrl, req.GithubUrl, req.ID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
