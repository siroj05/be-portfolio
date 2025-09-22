package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func ProfileRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewProfileRepository(db)
	handler := handlers.NewProfileHandler(repo)

	r.Handle("/profile/save", middleware.JWTauth(http.HandlerFunc(handler.CreateProfile))).Methods("POST")
}
