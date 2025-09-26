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

func ProjectsRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewProjectRepository(db)
	handler := handlers.NewProjectHandler(repo)

	file := middleware.NewRateLimiterStore(
		50*time.Millisecond, // tiap 50ms boleh 1 request
		50,                  // burst max 50
	)

	api := middleware.NewRateLimiterStore(
		200*time.Millisecond, // tiap 200ms boleh 1 request
		20,                   // burst max 20
	)

	fs := http.FileServer(http.Dir("./uploads"))
	// r.PathPrefix("/uploads/").Handler(file.RateLimitMiddleware(http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads")))))
	r.PathPrefix("/uploads/").Handler(file.RateLimitMiddleware(http.StripPrefix("/uploads", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=86400") // 1 hari
		fs.ServeHTTP(w, r)
	}))))

	r.Handle("/projects/save", middleware.JWTauth(http.HandlerFunc(handler.CreateProject))).Methods("POST")
	r.Handle("/projects/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.DeleteProject))).Methods("DELETE")
	r.Handle("/projects/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.GetProjectById))).Methods("GET")
	r.Handle("/projects/update", middleware.JWTauth(http.HandlerFunc(handler.UpdateProject))).Methods("PUT")

	r.Handle("/projects", api.RateLimitMiddleware(http.HandlerFunc(handler.GetAllProjects))).Methods("GET")
}
