# Transcribe Service

## Установка и запуск

### 1. Клонирование репозитория
```bash
git clone https://github.com/your-repo/transcribe_service.git
cd transcribe_service
```

### 2. Установка модели Vosk

Выберите подходящую модель в зависимости от точности и производительности:

#### 🔹 **Максимальная точность (большая модель)**
```bash
make install-max-model
```

#### 🔹 **Средняя точность (средняя модель)**
```bash
make install-medium-model
```

#### 🔹 **Минимальная точность (маленькая модель, быстрее работает)**
```bash
make install-small-model
```

### 3. Проверка установки модели
```bash
make check
```
**Ожидаемый результат:**
```
Модель Vosk найдена!
```

### 4. Сборка Docker-образа
```bash
make build
```

### 5. Запуск сервиса
```bash
make run
```
**Ожидаемый вывод:**
```
>> App is starting on :8080..
```

## 📡 Использование API

Отправьте аудиофайл для распознавания:
```bash
curl -X POST -F "audio=@/path/to/audiofile.wav" http://dockerhost:8080/api/v1/transcribe
```

Пример ответа:
```json
{
    "Text": "Пример распознанного текста"
}
```



