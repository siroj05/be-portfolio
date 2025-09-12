package dto

type SkillDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CategoriesDto struct {
	ID       string     `json:"id"`
	Category string     `json:"category"`
	Skills   []SkillDto `json:"skills"`
}
