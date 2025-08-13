package dto

type MessageDto struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}
