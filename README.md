# payment-system

Read Development instructions in [CONTRIBUTING.md](CONTRIBUTING.md)

## What has been done

The following decisions were taken
- Monorepo / Mono repository - a single repository to place the two services
- Single `go.mod` file

## Why

### Why Monorepo?

For simplicity. If it becomes complex to handle a monorepo, we can split into separate repositories

### Why single `go.mod` file?

For simplicity. If it becomes complex to handle a single `go.mod` file to manage the dependencies for both the services, then we can use separate `go.mod` files to manage the dependencies for both the services

## How to run it

### How to run it locally without using Containers or Kubernetes

Check [CONTRIBUTING.md](CONTRIBUTING.md)
