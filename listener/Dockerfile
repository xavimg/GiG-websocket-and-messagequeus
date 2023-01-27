FROM golang:1.19.2-alpine

WORKDIR /listener

# current directory to /app
COPY . .

RUN go build -o main listener/cmd/main.go

EXPOSE 8081

CMD ["/listener/main"]