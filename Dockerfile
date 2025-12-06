# -------- Build Stage --------
FROM golang:1.22 AS build

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the Go app (disable CGO, linux OS, production flags)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main

# -------- Runtime Stage --------
FROM alpine:latest

# Install any dependencies (curl, tzdata, etc.)
RUN apk add --no-cache curl tzdata

WORKDIR /root

# Copy built binary and .env file
COPY --from=build /app/main .
COPY --from=build /app/.env .

# Expose the port your app uses
EXPOSE 3000

# Run the app
CMD ["./main"]
