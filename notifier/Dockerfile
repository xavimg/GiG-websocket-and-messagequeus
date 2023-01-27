FROM golang:1.19.2-alpine

WORKDIR /notifier

# current directory to /app
COPY . .

RUN go build -o main notifier/cmd/main.go

EXPOSE 8080

CMD ["/notifier/main"]