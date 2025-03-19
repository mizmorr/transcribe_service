# === Этап 1: Сборка Go-приложения ===
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .  

ADD go.mod .
ADD go.sum .
RUN go mod download

RUN go build -o /cmd/app /cmd/main.go

# === Этап 2: Установка Python-приложения ===
FROM python:3.12-slim AS python_builder
WORKDIR /app
RUN pip install -r requirements.txt

# === Этап 3: Финальный минимальный образ ===
FROM debian:bullseye-slim

# Копируем Go-бинарник 
COPY --from=builder /cmd/app /usr/local/bin/app

# === Копирование модели отдельным шагом для кеширования ===
COPY cmd/model_ru/ /app/cmd/model_ru/

# Установка Python для запуска скрипта
RUN apt-get update && \
    apt-get install -y python3 && \
    rm -rf /var/lib/apt/lists/*

# Указываем рабочую директорию
WORKDIR /app

CMD ["/usr/local/bin/app"]
