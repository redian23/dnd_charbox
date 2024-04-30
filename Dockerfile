# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd/
RUN CGO_ENABLED=0 go build -o charbox_test

COPY  pkg/db/ /opt/charbox/database/

EXPOSE 4050
CMD ["./charbox"]
