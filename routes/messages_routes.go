package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func MessagesRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewMessagesRepository(db)
	handler := handlers.NewMessagesHandler(repo)

	r.HandleFunc("/messages/send", handler.CreateMessage).Methods("POST")

	// with middleware
	r.Handle("/messages", middleware.JWTauth(http.HandlerFunc(handler.GetAllMessages))).Methods("GET")
	r.Handle("/messages/{id:[0-9]+}", middleware.JWTauth(http.HandlerFunc(handler.DeleteMessages))).Methods("DELETE")
	r.Handle("/messages/delete-all", middleware.JWTauth(http.HandlerFunc(handler.DeleteAllMessages))).Methods("DELETE")
	r.Handle("/messages/{id}/mark", middleware.JWTauth(http.HandlerFunc(handler.MarkReadMessage))).Methods("PUT")
	r.Handle("/messages/mark-all", middleware.JWTauth(http.HandlerFunc(handler.MarkAllMessage))).Methods("PUT")
}
