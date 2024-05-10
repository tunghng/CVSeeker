# Use an official Golang runtime as a parent image
FROM golang:1.22.0 as builder

# Set the Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Set the Working Directory inside the container for building the project
WORKDIR /app/cmd/CVSeeker

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a smaller base image to run the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Install bash
RUN apk add --no-cache bash

# Add wait-for-it
COPY --from=builder /app/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

WORKDIR /root/

# Copy the built binary and any other necessary files
COPY --from=builder /app/cmd/CVSeeker/main .

COPY --from=builder /app/cmd/CVSeeker/statics /root/statics

# Run the binary
CMD ["./main"]
