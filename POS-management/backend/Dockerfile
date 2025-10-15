# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd/server

# Run stage
FROM alpine:3.20
WORKDIR /app
ENV PORT=8080
COPY --from=builder /app/server /app/server
EXPOSE 8080
CMD ["/app/server"]
