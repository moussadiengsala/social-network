# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Install SQLite
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Copy the entire directory into the container
COPY . .

# Build the Go application
RUN go build -o main server.go

# Expose the port that the app will run on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]