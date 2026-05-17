# --- STAGE 1: Build Stage ---
FROM golang:1.26.3-alpine AS builder

# Set build argument for the service to build. For example - payment-gateway, payment-processor)
ARG SERVICE_NAME

WORKDIR /app

# Copy dependency files first to leverage image layer caching
# in case there are no changes in the list of dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire monorepo to access any shared internal packages.
# We have .dockerignore file to take care of ignoring any unnecessary
# files
COPY . .

# Build a static binary for the specific service
RUN CGO_ENABLED=0 go build -o /service ./cmd/${SERVICE_NAME}

# --- STAGE 2: Final Production Stage ---
FROM scratch AS final

# Copy only the compiled binary from the builder stage
COPY --from=builder /service /service

EXPOSE 8080

ENTRYPOINT ["/service"]
