# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o portfolio-app main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/portfolio-app .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 8080

ENTRYPOINT ["/app/portfolio-app"]