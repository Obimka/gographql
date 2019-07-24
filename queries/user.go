package queries

import (
	"cadoles/graphql/postgres"
	"cadoles/graphql/types"
	"log"

	"github.com/graphql-go/graphql"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	log.Print("GetUserQuery")
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var user types.User
			users := []types.User{}

			sqlStatement := GET_USERS
			rows, err := postgres.DB.Query(sqlStatement)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {

				err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname)
				users = append(users, user)
				if err != nil {
					log.Fatalln(err)
				}
			}
			return users, nil
		},
	}
}
