package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/handlers"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/internal/repository"
)

func ProjectsRoutes(r *mux.Router, db *sql.DB) {
	repo := repository.NewProjectRepository(db)
	handler := handlers.NewProjectHandler(repo)

	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))
	r.Handle("/projects/save", middleware.JWTauth(http.HandlerFunc(handler.CreateProject))).Methods("POST")
	r.Handle("/projects", middleware.JWTauth(http.HandlerFunc(handler.GetAllProjects))).Methods("GET")
	r.Handle("/projects/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.DeleteProject))).Methods("DELETE")
	r.Handle("/projects/{id:[0-9a-fA-F\\-]{36}}", middleware.JWTauth(http.HandlerFunc(handler.GetProjectById))).Methods("GET")
}
