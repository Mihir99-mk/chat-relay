# syntax=docker/dockerfile:1

# ===== Build Stage =====
FROM golang:1.24.1 as builder

WORKDIR /app

# Cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

# ===== Runtime Stage =====
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# Copy binary
COPY --from=builder /app/auth .

# Run it
ENTRYPOINT ["/app/auth"]
