# gographql

Exemple d'API GraphQL, PostgresQL, en Go

```
go build && go run .
```

Accéder à l'API :
```
http://localhost:8383/graphql
```

Accéder à l'UI graphiql :
```
http://localhost:8383
```

## Query
Toute rếquete visant à récupérer une donnée est appellée une 'Query'

```
# Query via curl

curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"query": " query { user { id,firstname,lastname,roles{name} } }"}' \
    http://localhost:8383/graphql

```
```
# Query via GraphQL

query GetAllUsers {
  user {
    id
    firstname
    lastname
    roles {
      name
    }
  }
}

```

## Mutation
Toute requếte visant à modifier une donnée est appellée une 'Mutation'

```
# Mutation via curl

curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"query": "mutation { createUser(firstname: \"John\", lastname: \"Snow\") { id,firstname,lastname } }"}' \
    http://localhost:8383/graphql

```

```
# Mutation via GraphQL

mutation CreateUser($firstname: String!, $lastname: String!) {
  createUser(firstname: $firstname, lastname: $lastname) {
    id
    firstname
    lastname
  }
}

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

psql -U postgres -h localhost
> create user graphql;
> alter user graphql with encrypted password 'graphql';
> create database graphql;
> GRANT ALL PRIVILEGES ON DATABASE "graphql" to graphql;
> \exit

psql -U graphql -h localhost
> CREATE TABLE users (
id SERIAL PRIMARY KEY,
firstname TEXT,
lastname TEXT
  );

> CREATE TABLE roles (
id SERIAL PRIMARY KEY,
name TEXT 
  );

```
