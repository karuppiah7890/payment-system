# Contributing to the services

## How to run it locally without using Containers or Kubernetes

### Building the services from source

This repository is a mono repository. It contains two services - `payment-gateway` and `payment-processor`

The two services have been built and tested with Go (Golang) version `1.26.3`

Please download and install Golang from https://go.dev/dl - preferably version `1.26.3` or newer but with same major version `1` - with newer minor and/or patch version

You can build the two services by simply running

```bash
go build -v ./cmd/payment-gateway

go build -v ./cmd/payment-processor
```

### Running the services

After building it from source, you can run the services by simply running

```bash
./payment-gateway
```

```bash
./payment-processor
```

## How to run it locally with just Containers

There is a generic `Dockerfile`s present for both the services

You can build the container images for the two services using a container build tool like `docker`. You can also choose other container build tools like `podman` etc

An example using `docker` CLI and Docker daemon -

Build `payment-gateway` like this -

```bash
docker build --build-arg SERVICE_NAME=payment-gateway -t payment-gateway .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --build-arg SERVICE_NAME=payment-gateway -t payment-gateway .
```

Run `payment-gateway` like this -

```bash
docker run --rm --publish 8080:8080 payment-gateway
```

Build `payment-processor` like this -

```bash
docker build --build-arg SERVICE_NAME=payment-processor -t payment-processor .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --build-arg SERVICE_NAME=payment-processor -t payment-processor .
```

Run `payment-processor` like this -

```bash
docker run --rm --publish 8080:8080 payment-processor
```

> [!NOTE]
> If you want to run both the services using containers and also expose the ports of both services to the host, then you need to choose different ports on the host as in the container level - both services expose 8080 and you can't map both the 8080 ports to same 8080 port on the host

## Debugging Docker builds

To debug what's being sent as part of the build context, please use the `Dockerfile.debug` file like this -

```bash
DOCKER_BUILDKIT=0 docker build -f Dockerfile.debug
```

Accordingly modify the `.dockerignore` file
