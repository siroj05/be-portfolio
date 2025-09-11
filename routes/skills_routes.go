package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func SkillsRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewSkillsRepository(db)
	handler := handlers.NewSkillsHandler(repo)

	r.Handle("/skills/save", middleware.JWTauth(http.HandlerFunc(handler.CreateSkill))).Methods("POST")
}
