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
docker build --build-arg SERVICE_NAME=payment-gateway --tag payment-gateway .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --build-arg SERVICE_NAME=payment-gateway --tag payment-gateway .
```

And if the build is cached, use below to create a new build without cache

```bash
docker build --no-cache --build-arg SERVICE_NAME=payment-gateway --tag payment-gateway .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --no-cache --build-arg SERVICE_NAME=payment-gateway --tag payment-gateway .
```


Run `payment-gateway` like this -

```bash
docker run --rm --publish 8080:8080 payment-gateway
```

Build `payment-processor` like this -

```bash
docker build --build-arg SERVICE_NAME=payment-processor --tag payment-processor .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --build-arg SERVICE_NAME=payment-processor --tag payment-processor .
```

And if the build is cached, use below to create a new build without cache

```bash
docker build --no-cache --build-arg SERVICE_NAME=payment-processor --tag payment-processor .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --no-cache --build-arg SERVICE_NAME=payment-processor --tag payment-processor .
```

Run `payment-processor` like this -

```bash
docker run --rm --publish 8080:8080 payment-processor
```

> [!NOTE]
> If you want to run both the services using containers and also expose the ports of both services to the host, then you need to choose different ports on the host as in the container level - both services expose 8080 and you can't map both the 8080 ports to same 8080 port on the host

> [!NOTE]
> If you want to run both the services using containers, you can use Docker Compose. Read the below section on using Docker Compose

## How to run it locally with just Containers using Docker Compose

The below command will build the two images and also run them for you

```bash
docker compose up --build --detach
```

The Docker Compose file is defined in such a way that both the services are accessible to each other through the default network. And the payment-processor service is not exposed to the outside world, in this case, the host, and only the payment-gateway is exposed to the host through 8080 port and is accessible only through the loopback network interface - `localhost` or `127.0.0.1` and not all network interfaces for security reasons - so as to not expose the service to outside devices that can access the host through other network interfaces

To run the containers without building the image, just run

```bash
docker compose up --detach
```

If you want to just build the two images, just run

```bash
docker compose build
```

And if the build is cached, use below to create a new build without cache

```bash
docker compose build --no-cache
```

## Debugging Docker builds

To debug what's being sent as part of the build context, please use the `Dockerfile.debug` file like this -

```bash
DOCKER_BUILDKIT=0 docker build --file Dockerfile.debug
```

Accordingly modify the `.dockerignore` file

And if the build is cached, use below to create a new build without cache

```bash
DOCKER_BUILDKIT=0 docker build --no-cache --file Dockerfile.debug
```

## How to run it locally with just Containers using Kubernetes

We'll be using `helm` tool to deploy (install) and manage our services. Management means - get deployment information, upgrade our services, delete our services

Please install `helm` by following the official Helm website https://helm.sh/docs/intro/install or from the official Helm releases - https://github.com/helm/helm/releases

Once installed, also check if a tool like `minikube` or `kind` or similar is installed to run local Kubernetes clusters

We'll be using `minikube` with a driver like `docker` for example

```bash
minikube start
```

And once the Kubernetes cluster is ready, ensure that the container images are available in the worker node's container runtime. For example, for `minikube`, you can access the container runtime like this -

First get details to connect to the container daemon

```bash
minikube docker-env
```

For a specific minikube profile, you can do this -

```bash
minikube --profile <profile-name> docker-env
```

Then run it the commands that it gives in your shell. You can also do this -

```bash
eval $(minikube docker-env)
```

Once you do this, you can see the list of containers running in the daemon like this -

```bash
docker ps
```

Now just build the images using `docker build` or `docker compose build`. We'll use `docker compose build` to build both the images together, like this -

```bash
docker compose build
```

Once done, check if the container daemon has the images of `payment-gateway` and `payment-processor` services using this -

```bash
docker images
```

You should see `payment-gateway:latest`, `payment-processor:latest`

Now you can use the helm chart to run the services

```bash
helm install payment-gateway helm-chart --set image.repository=payment-gateway --set image.tag=latest --set service.port=8080 --set livenessProbe.httpGet.path=/healthz --set readinessProbe.httpGet.path=/healthz

helm install payment-processor helm-chart --set image.repository=payment-processor --set image.tag=latest --set service.port=8080 --set livenessProbe.httpGet.path=/healthz --set readinessProbe.httpGet.path=/healthz
```

To run it easily with lesser command line arguments, you can use the helm values yaml files for the two services like this -

```bash
helm install payment-gateway helm-chart --values payment-gateway-helm-values.yaml

helm install payment-processor helm-chart --values payment-processor-helm-values.yaml
```
