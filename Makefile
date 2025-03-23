install-max-model:
	wget https://alphacephei.com/vosk/models/vosk-model-ru-0.10.zip
	unzip vosk-model-ru-0.10.zip
	mv vosk-model-ru-0.10 model_ru
	rm vosk-model-ru-0.10.zip

install-medium-model:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.42.zip
	unzip vosk-model-small-ru-0.42.zip
	mv vosk-model-small-ru-0.42 model_ru
	rm vosk-model-small-ru-0.42.zip

install-small-model:
	wget https://alphacephei.com/vosk/models/vosk-model-ru-0.22.zip
	unzip vosk-model-ru-0.22.zip
	mv vosk-model-ru-0.22 model_ru
	rm vosk-model-ru-0.22.zip

check:
	@if [ -d "model_ru" ]; then \
			echo "Модель Vosk найдена!"; \
	else \
			echo "Модель Vosk не найдена.."; \
	fi

build:
	@echo ">> App is building .."
	docker build -t transcriber:custom .

run:
	@if docker images --format "{{.Repository}}:{{.Tag}}" | grep -q "^transcribe:custom$$"; then \
		@echo ">> App is starting on :8080.."
		docker run -p 8080:8080 transcriber:v1
	else \
		echo "Образ не найден. Run make build"; \
	fi

