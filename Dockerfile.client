# Dockerfile.client
FROM golang:1.18 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o client client.go

# Use a smaller base image for the final container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/client .

CMD ["./client"]
