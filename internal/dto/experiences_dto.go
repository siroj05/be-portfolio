package dto

type ExperiencesDto struct {
	ID          string `json:"id"`
	Office      string `json:"office"`
	Position    string `json:"position"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
	Present     bool   `json:"present"`
}

type ExperiencesListDto struct {
	ID          string `json:"id"`
	Office      string `json:"office"`
	Position    string `json:"position"`
	Start       string `json:"start"`
	End         string `json:"end,omitempty"`
	Description string `json:"description"`
	Present     bool   `json:"present"`
}
