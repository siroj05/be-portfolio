package models

type ProjectModel struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TechStack   string `json:"techStack"`
	DemoUrl     string `json:"demoUrl"`
	GithubUrl   string `json:"githubUrl"`
}
