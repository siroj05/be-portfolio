package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
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
		response.Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.Repo.Create(ctx, req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to create message", err.Error())
		return
	}

	response.Success(w, "Data berhasil diambil", req)
}

func (h *MessagesHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	result, err := h.Repo.GetAll(ctx)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to get message", err.Error())
		return
	}

	response.Success(w, "Data berhasil diambil", result)
}

func (h *MessagesHandler) DeleteMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	i64 := int64(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Invalid id", err.Error())
		return
	}

	err = h.Repo.Delete(ctx, i64)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete message", err.Error())
		return
	}

	response.Success(w, "Delete message successfully", nil)
}

func (h *MessagesHandler) MarkReadMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	i64 := int64(id)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Invalid id", err.Error())
		return
	}

	var req dto.MarkMessageDto
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err = h.Repo.Mark(ctx, i64, req)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to read message", err.Error())
		return
	}

	response.Success(w, "Success to read message", nil)
}
