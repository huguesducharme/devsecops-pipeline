# Стадия 1: Сборка (свежий Go, чтобы работали зависимости)
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# Стадия 2: Исполнение (старая ОС с уязвимостями для Trivy)
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
