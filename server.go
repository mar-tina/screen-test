package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/Salaton/screen-test.git/db"
	"github.com/Salaton/screen-test.git/graph"
	"github.com/Salaton/screen-test.git/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	if err := InitDB(); err != nil {
		log.Fatalf("FAILED TO CONNECT TO DB %v", err.Error())
	}

	// graph.DB.CreateCustomer()
	// graph.DB.CreateOrder()
	// graph.DB.FetchAllOrdersForCustomer("8888")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func InitDB() error {
	graph.DB = &db.PostgresClient{}
	dsn := "user=martina password=pass dbname=salatoons port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return graph.DB.Open(dsn)
}

