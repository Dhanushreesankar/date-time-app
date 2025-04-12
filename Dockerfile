# Use the official Golang image as the base
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy all Go files into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the host machine
EXPOSE 8080

# Start the Go app when the container launches
CMD ["./main"]
