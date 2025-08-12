package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
)

type MessagesHandler struct {
	Repo interfaces.MessagesRepository
}

func NewMessagesHandler(repo interfaces.MessagesRepository) *MessagesHandler {
	return &MessagesHandler{
		Repo: repo,
	}
}

func (h *MessagesHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	var req dto.MessageDto
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Repo.Create(ctx, req)

	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
