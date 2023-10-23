# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 3001
EXPOSE 3001

# Command to run the Go application
CMD ["./main"]