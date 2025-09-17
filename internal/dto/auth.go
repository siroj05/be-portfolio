package dto

type LoginDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GetMeDto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
