# Use golang:1.18 as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application
RUN go build -o farmze-backend-api .

# Expose port 8080 for the API
EXPOSE 8080

# Command to run the application
CMD ["./farmze-backend-api"]