package dto

type SkillDto struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CategoriesDto struct {
	Category string     `json:"category"`
	Skills   []SkillDto `json:"skills"`
}
