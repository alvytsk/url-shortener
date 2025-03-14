# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS build

# Установка зависимостей для CGO и SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Установка swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/app/main.go

# Включаем CGO и собираем бинарник
RUN CGO_ENABLED=1 go build -o link-shortener ./cmd/app/main.go

FROM alpine:latest

# Важно! SQLite runtime-зависимости
RUN apk add --no-cache sqlite-libs

WORKDIR /app
COPY --from=build /app/link-shortener /app/link-shortener
COPY --from=build /app/docs /app/docs

EXPOSE 8080

CMD ["/app/link-shortener"]
