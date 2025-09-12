package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/siroj05/portfolio/internal/dto"
)

type SkillsRepository struct {
	db *sql.DB
}

func NewSkillsRepository(db *sql.DB) *SkillsRepository {
	return &SkillsRepository{db: db}
}

func (r *SkillsRepository) Create(ctx context.Context, req dto.CategoriesDto) error {
	CategoryID := uuid.New().String()
	_, err := r.db.ExecContext(ctx, "INSERT INTO categories (id, name) VALUES (?, ?)", CategoryID, req.Category)
	if err != nil {
		return err
	}

	for _, skill := range req.Skills {
		skillID := uuid.New().String()
		_, err := r.db.ExecContext(ctx, "INSERT INTO skills (id, category_id, name, icon) VALUES (?, ?, ?, ?)", skillID, CategoryID, skill.Name, skill.Icon)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SkillsRepository) GetAll(ctx context.Context) ([]dto.CategoriesDto, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT 
			c.id AS category_id,
			c.name AS category_name,
			s.id AS skill_id,
			s.name AS skill_name,
			s.icon AS skill_icon
		FROM categories c
		LEFT JOIN skills s ON c.id = s.category_id
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoriesMap := make(map[string]*dto.CategoriesDto)
	BASE_URL := "http://localhost:8080/"
	for rows.Next() {
		var (
			categoryID   string
			categoryName string
			skillID      string
			skillName    string
			skillIcon    string
		)

		err := rows.Scan(&categoryID, &categoryName, &skillID, &skillName, &skillIcon)
		if err != nil {
			return nil, err
		}

		if _, ok := categoriesMap[categoryID]; !ok {
			categoriesMap[categoryID] = &dto.CategoriesDto{
				ID:       categoryID,
				Category: categoryName,
				Skills:   []dto.SkillDto{},
			}
		}

		if skillID != "" {
			categoriesMap[categoryID].Skills = append(categoriesMap[categoryID].Skills, dto.SkillDto{
				ID:   skillID,
				Name: skillName,
				Icon: BASE_URL + skillIcon,
			})
		}
	}

	result := make([]dto.CategoriesDto, 0, len(categoriesMap))
	for _, cat := range categoriesMap {
		result = append(result, *cat)
	}

	return result, nil
}

func (r *SkillsRepository) GetSkillsByCategory(ctx context.Context, id string) ([]dto.SkillDto, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT icon FROM skills WHERE category_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []dto.SkillDto
	for rows.Next() {
		var s dto.SkillDto
		rows.Scan(&s.Icon)
		skills = append(skills, s)
	}

	return skills, nil
}

func (r *SkillsRepository) Delete(ctx context.Context, id string) error {
	// begin tx nih penting jadi kalau ada operasi query yang ada relasinya, terus gagal dia bakal otomatis di rollback
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// hapus dulu skill nya
	_, err = tx.ExecContext(ctx, "DELETE FROM skills WHERE category_id = ?", id)
	if err != nil {
		// kalo operasi gagal ya di rollback
		tx.Rollback()
		return err
	}

	// hapus kategory
	_, err = tx.ExecContext(ctx, "DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
