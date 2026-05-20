// @title Portfolio API
// @version 1.0
// @description API Server for Portfolio Website.
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/siroj05/portfolio/config"
	_ "github.com/siroj05/portfolio/docs"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.LoadEnv()
	config.LoadImgUrl()
	config.GetConnection()
	defer config.DB.Close()

	r := mux.NewRouter()

	// Swagger UI route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// routes API endpoint
	routes.MessagesRoutes(r, config.DB)
	routes.AuthRoutes(r, config.DB)
	routes.ExperiencesRoutes(r, config.DB)
	routes.ProjectsRoutes(r, config.DB)
	routes.SkillsRoutes(r, config.DB)
	routes.ProfileRoutes(r, config.DB)
	routes.HealthRoutes(r, config.DB)
	// handle with middleware
	handlerWithMiddleware := middleware.Logging(r)

	// cors
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	finalHandler := corsHandler.Handler(handlerWithMiddleware)

	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", finalHandler)
	error := http.ListenAndServe(":8080", finalHandler)
	if error != nil {
		log.Fatal(error)
	}
}
