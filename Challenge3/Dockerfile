# Use the official Golang image as a parent image
FROM golang:1.20

# Set the current working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Check if go.mod already exists
RUN if [ ! -f "go.mod" ]; then go mod init my-go-app; fi

# Tidy up the Go module if go.mod exists
RUN if [ -f "go.mod" ]; then go mod tidy; fi

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]

