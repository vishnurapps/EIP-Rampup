# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Vishnu R<vishnurapps@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

#to get mod and sum files
RUN go mod init vishnu

# copying server.go into app which is working directory
COPY server.go .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download


# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 9091

# Command to run the executable
CMD ["./main"]

