# Стадия 1: Сборка приложения
FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
# Собираем бинарник
RUN go build -o main .

# Стадия 2: Финальный образ (намеренно берем старую версию с CVE)
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
