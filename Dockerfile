# Стадия 1: Сборка приложения
FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .

# Синхронизируем и скачиваем зависимости внутри контейнера
RUN go mod tidy

# Собираем бинарник
RUN go build -o main .
