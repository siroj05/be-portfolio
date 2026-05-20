# Stage 1: Build the Go binary
FROM golang:alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the Go application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/portfolio-be .

# Stage 2: Final minimal image
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/portfolio-be /app/portfolio-be

# Copy uploads folder
COPY --from=builder /app/uploads /app/uploads

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["/app/portfolio-be"]
