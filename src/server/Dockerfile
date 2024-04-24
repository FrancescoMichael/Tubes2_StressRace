FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "cmd/app/main.go"]