# Use the latest Golang image as the base image
FROM golang:latest

# Set the working directory inside the container to /app
WORKDIR /

# Copy the current directory contents into the container at /
COPY . .

# Download dependencies listed in go.mod and go.sum
RUN go mod download

# Build the Go application and name the output binary as 'main'
RUN go build -o main .

# Command to run the executable when the container starts
CMD ["./main"]
