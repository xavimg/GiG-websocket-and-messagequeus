FROM golang:1.19.2-alpine AS builder

WORKDIR /notifier

# current directory to /app
COPY . .

RUN go build -o main cmd/main.go

# Run stage
FROM alpine
WORKDIR /notifier
COPY --from=builder /notifier/main .

EXPOSE 8080

CMD ["/notifier/main"]