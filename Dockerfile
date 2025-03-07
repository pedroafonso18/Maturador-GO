# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application from the cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o maturador-go ./cmd/

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/maturador-go /app/
# Copy resources directory
COPY --from=builder /app/resources /app/resources

# Set the working directory to be the same as in the code
ENV PATH="/app:${PATH}"

# Run as non-privileged user
RUN adduser -D appuser
USER appuser

CMD ["/app/maturador-go"]
