# Usar a imagem oficial do Golang
FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]