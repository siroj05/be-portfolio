package dto

import "time"

type ExperiencesDto struct {
	// ID          string    `json:"id"`
	Office      string    `json:"office"`
	Position    string    `json:"position"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Description string    `json:"description"`
}
