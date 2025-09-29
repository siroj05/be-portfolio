package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/dto"
	"github.com/siroj05/portfolio/internal/repository/interfaces"
	"github.com/siroj05/portfolio/internal/response"
	"github.com/siroj05/portfolio/utils"
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
	var req dto.CreateMessageDto
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// verifikasi token turnstile
	ok, err := utils.VerifyTurnstile(req.Token)
	if err != nil || !ok {

		msg := "Invalid captcha"
		detail := ""
		if err != nil {
			detail = err.Error()
			msg = "Captcha error"
		}

		response.Error(w, http.StatusBadRequest, msg, detail)
		return
	}

	// Validasi
	if len(req.Email) > 50 {
		response.Error(w, http.StatusBadRequest, "Email too long (max 50)", err.Error())
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		response.Error(w, http.StatusBadRequest, "Email is required", err.Error())
		return
	}

	if strings.TrimSpace(req.Message) == "" {
		response.Error(w, http.StatusBadRequest, "Message is required", err.Error())
		return
	}

	if len(req.Message) > 500 {
		response.Error(w, http.StatusBadRequest, "Message too long (max 500)", err.Error())
		return
	}

	// simpan ke db
	err = h.Repo.Create(ctx, req)

	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to create message", err.Error())
		return
	}

	response.Success(w, "Create message success", req)
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

	response.Success(w, "Successfully get all massage", result)
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

	response.Success(w, "Successfully deleted message", nil)
}

func (h *MessagesHandler) DeleteAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	err := h.Repo.DeleteAll(ctx)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to delete all messages", err.Error())
		return
	}

	response.Success(w, "Successfully deleted all messages", nil)
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

func (h *MessagesHandler) MarkAllMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	err := h.Repo.MarkAll(ctx)
	if err != nil {
		log.Println(err)
		response.Error(w, http.StatusInternalServerError, "Failed to mark all messages as read", err.Error())
		return
	}

	response.Success(w, "Successfully marked all messages as read", nil)
}
