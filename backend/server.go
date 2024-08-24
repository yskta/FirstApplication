package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/yskta/taskmaster/backend/database"
	"github.com/yskta/taskmaster/backend/graph"
	// "github.com/yskta/taskmaster/backend/graph/generated"
)

const defaultPort = "8080"

func runServer() {
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

	//データベースにテストデータを投入する。
	if err := database.SeedData(db); err != nil {
		log.Fatalf("failed to seed data: %v", err)
	}

	// Create a new resolver with the database connection
	resolver := &graph.Resolver{DB: db}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
