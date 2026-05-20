package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siroj05/portfolio/internal/response"
)

// HealthRoutes registers healthcheck endpoints
func HealthRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if db != nil {
			if err := db.Ping(); err != nil {
				response.Error(w, http.StatusInternalServerError, "database disconnected", err.Error())
				return
			}
		}
		response.Success(w, "healthy ok", map[string]string{
			"database": "connected",
			"status":   "healthy",
		})
	}).Methods("GET")

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if db != nil {
			if err := db.Ping(); err != nil {
				response.Error(w, http.StatusInternalServerError, "database disconnected", err.Error())
				return
			}
		}
		response.Success(w, "healthy ok", map[string]string{
			"database": "connected",
			"status":   "healthy",
		})
	}).Methods("GET")
}
