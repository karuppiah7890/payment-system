# Contributing to the services

## Building the services from source

This repository is a mono repository. It contains two services - `payment-gateway` and `payment-processor`

The two services have been built and tested with Go (Golang) version `1.26.3`

Please download and install Golang from https://go.dev/dl - preferably version `1.26.3` or newer but with same major version `1` - with newer minor and/or patch version

You can build the two services by simply running

```bash
go build -v ./cmd/payment-gateway

go build -v ./cmd/payment-processor
```

## Running the services

After building it from source, you can run the services by simply running

```bash
./payment-gateway
```

```bash
./payment-processor
```
