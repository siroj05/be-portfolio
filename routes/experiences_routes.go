package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func ExperiencesRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewExperiencesRepository(db)
	handler := handlers.NewExperiencesHandler(repo)

	r.Handle("/experiences/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.GetExperienceById))).Methods("GET")
	r.Handle("/experiences/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.DeleteExperiences))).Methods("DELETE")
	r.Handle("/experiences/save", middleware.JWTauth(http.HandlerFunc(handler.CreateExperience))).Methods("POST")
	r.Handle("/experiences", middleware.JWTauth(http.HandlerFunc(handler.GetAllExperiences))).Methods("GET")
}
