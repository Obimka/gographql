# gographql

Exemple d'API GraphQL, PostgresQL, en Go

```
go build && go run .
```

## Query
Toute rếquete visant à récupérer une donnée est appellée une 'Query'

```
# Exemple de Query

curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"query": "{ user { id,firstname,lastname,roles{name} } }"}' \
    http://localhost:8383/

```

## Mutation
Toute requếte visant à modifier une donnée est appellée une 'Mutation'

```
curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"query": "mutation { createUser(firstname: \"John\", lastname: \"Snow\") { id,firstname,lastname } }"}' \
    http://localhost:8383

```