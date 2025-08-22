package dto

import (
	"github.com/siroj05/portfolio/utils"
)

type ExperiencesDto struct {
	// ID          string    `json:"id"`
	Office      string         `json:"office"`
	Position    string         `json:"position"`
	Start       utils.DateOnly `json:"start"`
	End         utils.DateOnly `json:"end"`
	Description string         `json:"description"`
}
