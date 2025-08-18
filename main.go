package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/siroj05/portfolio/config"
	"github.com/siroj05/portfolio/routes"
)

func main() {
	config.LoadEnv()
	config.GetConnection()
	defer config.DB.Close()

	r := mux.NewRouter()
	routes.MessagesRoutes(r, config.DB)
	routes.AuthRoutes(r, config.DB)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", corsHandler.Handler(r))
	error := http.ListenAndServe(":8080", corsHandler.Handler(r))
	if error != nil {
		log.Fatal(error)
	}
}
