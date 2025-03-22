# === Этап 1: Сборка Go-приложения ===
FROM golang:1.24 AS builder

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app ./cmd/main.go 

# === Этап 2: Установка Python-приложения ===

FROM debian:bullseye-slim
# Копируем Go-бинарник 

WORKDIR /app

COPY output.wav output.wav

# Копируем только скомпилированный Go-бинарник
COPY --from=builder /app/app ./app

# Копируем только зависимости Python
COPY src/requirements.txt ./requirements.txt

COPY model_ru/ ./model_ru/

COPY src/internal/transcriber/transcribe.py transcribe.py 


RUN apt-get update && \
    apt-get install -y python3 python3-pip ffmpeg && \
    pip3 install --no-cache-dir -r ./requirements.txt && \
    rm -rf /var/lib/apt/lists/*

# Запуск приложения
CMD ["./app"]