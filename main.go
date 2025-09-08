package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/siroj05/portfolio/config"
	"github.com/siroj05/portfolio/internal/middleware"
	"github.com/siroj05/portfolio/routes"
)

func main() {
	config.LoadEnv()
	config.GetConnection()
	defer config.DB.Close()

	r := mux.NewRouter()
	routes.MessagesRoutes(r, config.DB)
	routes.AuthRoutes(r, config.DB)
	routes.ExperiencesRoutes(r, config.DB)
	routes.ProjectsRoutes(r, config.DB)
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
