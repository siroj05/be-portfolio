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

func SkillsRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewSkillsRepository(db)
	handler := handlers.NewSkillsHandler(repo)

	// bikin limiter khusus untuk skill API
	// misal 1 request / 200ms (≈ 5 req/detik) dengan burst 20
	skillsLimiter := middleware.NewRateLimiterStore(
		200*time.Millisecond, // tiap 200ms boleh 1 request
		20,                   // burst max 20
	)

	// pake jwt auth
	r.Handle("/skills/save", middleware.JWTauth(http.HandlerFunc(handler.CreateSkill))).Methods("POST")
	r.Handle("/skills/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.DeleteCategory))).Methods("DELETE")

	// pake rate limit
	r.Handle("/skills", skillsLimiter.RateLimitMiddleware(http.HandlerFunc(handler.GetAllSkills))).Methods("GET")
}
