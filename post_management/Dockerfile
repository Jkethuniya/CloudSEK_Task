# syntax=docker/dockerfile:1
FROM golang:1.24.1 AS build


# Install dependencies
RUN apt-get update && apt-get install -y git


# Set working directory
WORKDIR /app

# Copy go mod and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source files
COPY . .

# Build the Go app
RUN go build -o main .

# Run the binary
CMD ["./main"]

