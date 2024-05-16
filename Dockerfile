# Builder stage
# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Download and cache dependencies
RUN go mod tidy

# Build
RUN go build -o main ./cmd/main.go



# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

# Expose port 8080 for the container
EXPOSE 8080


# Set the entry point of the container to the Go app executable
CMD ["/app/main"]