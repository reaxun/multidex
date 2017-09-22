# Multidex

[![Build Status](https://travis-ci.org/reaxun/multidex.svg?branch=master)](https://travis-ci.org/reaxun/multidex)


Pokemon-related APIs and applications.

## Starting the API

### Running the app locally

To start up the API on your local machine, run `go build -ldflags -s && ./multidex`.
The endpoints can then be reached at `localhost:12345`.

### Running the app in a container

This app can be built from its Dockerfile.

```
make container
docker run --publish 12345:12345 --name multidex --rm multidex
```

As above, the app can then be reached at `localhost:12345`.

### Running the app in kubernetes

This app can be run in a kubernetes cluster.

```
kubectl run multidex --replicas=3 --image=reaxun/multidex:latest --port=12345
kubectl expose deployment multidex --type=LoadBalancer --name=multidex

EXTERNAL_IP=$(kubectl get service multidex | awk '/multidex/ {print $3}')
PORT=$(kubectl get service multidex | awk '/multidex/ {print $4}' | cut -d ':' -f 2 | cut -d '/' -f 1)
```

The app can be reached at `http://$EXTERNAL_IP:$PORT`.

## Using the API

This application only supports GET requests, and can be used to retrieve information about pokemon, attacks, and types.
The following endpoints are available:

| Endpoint               | Effect                              |
|------------------------|-------------------------------------|
| `/pokemon`             | List all Pokemon                    |
| `/pokemon/{name}`      | Display a specific Pokemon by name  |
| `/pokemon/type/{type}` | List all Pokemon of a specific Type |
| `/attacks`             | List all Attacks                    |
| `/attacks/{name}`      | Display a specific Attack by name   |
| `/attacks/type/{type}` | List all Attacks of a specific Type |
| `/types`               | List all Types                      |
