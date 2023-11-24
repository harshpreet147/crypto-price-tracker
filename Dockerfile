# First stage: build the binary
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download the dependencies
COPY go.mod go.mod ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the binary with all necessary compiler flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage: create the runtime image
FROM alpine:latest

RUN apk add --no-cache tzdata

# Set the working directory and copy the binary from the builder stage
WORKDIR /root/
COPY --from=builder /app/main .
    
# Expose the port the server listens on
EXPOSE 7575

# Run the compiled binary
CMD ["./main"]
