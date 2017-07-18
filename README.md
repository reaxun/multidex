# Multidex

Pokemon-related APIs and applications.

## Starting the API

### Running the app locally

To start up the API on your local machine, run `go build -ldflags -s && ./multidex`.
The endpoints can then be reached at `localhost:12345`.

### Running the app in a container

This app can be built from its Dockerfile.

```
docker build -t multidex .
docker run --publish 12345:12345 --name multidex --rm multidex
```

## Using the API

This application only supports GET requests, and can be used to retrieve information about pokemon, attacks, and types.
The following endpoints are available:

`/pokemon`

List all Pokemon

`/pokemon/{name}`

Display a specific Pokemon by name

`/pokemon/type/{type}`

List all Pokemon of a specific Type

`/attacks`

List all Attacks

`/attacks/{name}'

Display a specific Attack by name

`/attacks/type/{type}`

List all Attacks of a specific Type

`/types`

List all Types
