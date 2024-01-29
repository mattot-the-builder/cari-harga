FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o bin/main cmd/app/main.go

CMD ["./bin/main"]
