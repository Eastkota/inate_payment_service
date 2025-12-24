# Use the official Go image with necessary tools for development
FROM golang:1.24-alpine

# Set environment variables for x86_64 (amd64) architecture
ENV GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory
COPY . .

# Build Service
RUN go build -o inatepaymentservice


# Expose the necessary port
EXPOSE 8103

# Command to run the service
CMD ["./inatepaymentservice"]

