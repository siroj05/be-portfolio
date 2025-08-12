package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/repository"
)

func MessagesRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewMessagesRepository(db)
	handler := handlers.NewMessagesHandler(repo)

	r.HandleFunc("/messages/send", handler.CreateMessage).Methods("POST")
}
