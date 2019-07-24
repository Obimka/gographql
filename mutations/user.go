package mutations

import (
	"cadoles/graphql/postgres"
	"cadoles/graphql/types"

	"github.com/graphql-go/graphql"
)

// GetCreateUserMutation creates a new user and returns it.
func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user := &types.User{
				Firstname: params.Args["firstname"].(string),
				Lastname:  params.Args["lastname"].(string),
			}

			_, err := postgres.DB.Exec(ADD_USER, user.Firstname, user.Lastname)
			if err != nil {
				panic(err)
			}
			return user, nil
		},
	}
}
