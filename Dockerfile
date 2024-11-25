# Use Go base image
FROM golang:1.20-alpine

# Set working directory in the container
WORKDIR /app

# Copy all files into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o weather-api .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./weather-api"]
