package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
)



func main() {
	const port = ":8082"

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}}