package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString:= os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not defined")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthcheck", routeHandler)
	apiRouter.Get("/err", errorRouteHandler)

	router.Mount("/v1", apiRouter)

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	fmt.Println("Server starting on port " + portString)

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal("Error running servre")
	}

	fmt.Println("Server running on port", portString)
}
