FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy everything first
COPY . .

# THEN resolve dependencies
RUN go mod tidy

# Build
RUN go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]