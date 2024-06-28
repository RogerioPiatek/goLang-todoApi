# Start from a base image with Go installed
FROM golang:1.22.4-bookworm

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

#Copy config.toml file
COPY config-docker.toml ./

# Download dependencies
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 9000 to the outside world
EXPOSE 9000

# Command to run the executable
CMD ["./main"]
