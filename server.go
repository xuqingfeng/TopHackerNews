package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/xuqingfeng/TopHackerNews/graph"
)

const defaultPort = "8080"
const defaultAppEnv = "prod"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = defaultAppEnv
	}

	router := chi.NewRouter()

	if appEnv == "local" {
		router.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:3000"},
		}).Handler)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
