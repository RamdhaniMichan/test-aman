# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8081

CMD ["./app"]