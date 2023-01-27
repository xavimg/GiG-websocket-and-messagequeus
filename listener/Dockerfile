FROM golang:1.19.2-alpine AS builder

WORKDIR /listener

# current directory to /app
COPY . .

RUN go build -o main cmd/main.go

# Run stage
FROM alpine
WORKDIR /listener
COPY --from=builder /listener/main .

EXPOSE 8081

CMD ["/listener/main"]