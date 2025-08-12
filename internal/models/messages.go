package models

type Message struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Messages  string `json:"messages"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}
