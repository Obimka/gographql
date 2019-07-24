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

## Postgres
Si besoin, lancer un serveur postgres via Docker 

```
mkdir $HOME/docker/postgres/volumes
docker run --rm -d \
     --name postgres \
     -e POSTGRES_PASSWORD=postgres \
     -p 5432:5432 \
     -v $HOME/docker/postgres/volumes:/var/lib/postgresql/data \
     postgres
```
