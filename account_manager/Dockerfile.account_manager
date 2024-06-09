# Dockerfile.account_manager
FROM golang:1.18 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o account_manager account_manager.go

# Use a smaller base image for the final container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/account_manager .

CMD ["./account_manager"]
