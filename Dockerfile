FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pkg/ ./pkg/

# Build the application
RUN go build -o main ./pkg/service/main.go

# Use a lightweight base image for the final stage
FROM alpine:latest

RUN addgroup --system appgroup && adduser --system appuser --ingroup appgroup
USER appuser

# Set the working directory
WORKDIR /app/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the default port
EXPOSE 8000

# Command to run the application
CMD ["/app/main"]
