package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/yskta/taskmaster/database"
	"github.com/yskta/taskmaster/graph"
	"github.com/yskta/taskmaster/graph/generated"
)

const defaultPort = "8080"

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new resolver with the database connection
	resolver := &graph.Resolver{DB: db}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
