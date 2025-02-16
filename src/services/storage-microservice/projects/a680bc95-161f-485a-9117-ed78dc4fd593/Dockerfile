FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod  ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the port the app runs on
EXPOSE 8080

# Run the server binary
CMD ["./server"]
