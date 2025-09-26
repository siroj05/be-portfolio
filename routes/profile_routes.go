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

func ProfileRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewProfileRepository(db)
	handler := handlers.NewProfileHandler(repo)

	api := middleware.NewRateLimiterStore(
		200*time.Millisecond, // tiap 200ms boleh 1 request
		20,                   // burst max 20
	)

	r.Handle("/profile/save", middleware.JWTauth(http.HandlerFunc(handler.CreateProfile))).Methods("POST")
	r.Handle("/profile/{id:[0-9]+}", middleware.JWTauth(http.HandlerFunc(handler.GetProfileById))).Methods("GET")

	r.Handle("/profile/me", api.RateLimitMiddleware(http.HandlerFunc(handler.GetProfile))).Methods("GET")
}
