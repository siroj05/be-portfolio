package dto

import (
	"github.com/siroj05/portfolio/utils"
)

type ExperiencesDto struct {
	ID          string         `json:"id"`
	Office      string         `json:"office"`
	Position    string         `json:"position"`
	Start       utils.DateOnly `json:"start"`
	End         utils.DateOnly `json:"end"`
	Description string         `json:"description"`
}

type ExperiencesListDto struct {
	ID          string `json:"id"`
	Office      string `json:"office"`
	Position    string `json:"position"`
	Start       string `json:"start"`
	End         string `json:"end,omitempty"`
	Description string `json:"description"`
}
