package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/PatriQ94/GoGraphQLAPITesting"
)

func main() {
	port := "8080"

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(GoGraphQLAPITesting.NewExecutableSchema(GoGraphQLAPITesting.Config{Resolvers: &GoGraphQLAPITesting.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
