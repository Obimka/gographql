package main

import (
	"log"
	"net/http"

	"cadoles/graphql/mutations"
	"cadoles/graphql/postgres"
	"cadoles/graphql/queries"
	"cadoles/graphql/security"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	graphiql "github.com/mnmtanish/go-graphiql"
)

func main() {

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	postgres.DBConnect()
	defer postgres.DBClose()

	http.Handle("/graphql", security.Handle(httpHandler))
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	log.Print("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}
