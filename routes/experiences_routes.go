package routes

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func ExperiencesRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewExperiencesRepository(db)
	handler := handlers.NewExperiencesHandler(repo)

	api := middleware.NewRateLimiterStore(
		200*time.Millisecond, // tiap 200ms boleh 1 request
		20,                   // burst max 20
	)

	r.Handle("/experiences/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.GetExperienceById))).Methods("GET")
	r.Handle("/experiences/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.DeleteExperiences))).Methods("DELETE")
	r.Handle("/experiences/save", middleware.JWTauth(http.HandlerFunc(handler.CreateExperience))).Methods("POST")
	r.Handle("/experiences/update", middleware.JWTauth(http.HandlerFunc(handler.UpadateExperience))).Methods("PUT")

	r.Handle("/experiences", api.RateLimitMiddleware(http.HandlerFunc(handler.GetAllExperiences))).Methods("GET")
}
