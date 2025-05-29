FROM golang:1.24-alpine AS builder

WORKDIR /app

# Add go.mod and go.sum before copying the rest to leverage Docker layer caching
COPY go.mod ./
RUN go mod download

# Copy the source code, including templates
COPY cmd/web/ ./cmd/web/

# Build the Go binary (make sure templates are in the right path for embedding)
WORKDIR /app/cmd/web
RUN go build -o /frontApp .

# Final minimal image
FROM alpine:latest

COPY --from=builder /frontApp /app/frontApp

CMD ["/app/frontApp"]