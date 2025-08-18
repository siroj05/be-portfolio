package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/repository"
)

func AuthRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewAuthRepository(db)
	handler := handlers.NewAuthHandler(repo)

	r.HandleFunc("/auth/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/auth/register", handler.CreateUser).Methods("POST")
	r.HandleFunc("/auth/logout", handler.LogoutUser).Methods("POST")
}
