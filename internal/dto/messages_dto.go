package dto

type MessageDto struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}

type MarkMessageDto struct {
	Mark bool `json:"mark"`
}

type CreateMessageDto struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}
