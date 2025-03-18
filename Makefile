install-max-model:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.42.zip
	unzip vosk-model-small-ru-0.42.zip
	mv vosk-model-ru-0.42 cmd/model_ru

install-small-model:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.22.zip
	unzip vosk-model-small-ru-0.22.zip
	mv vosk-model-ru-0.22 cmd/model_ru

check:
	@if [ -d "cmd/model_ru" ]; then \
			echo "Модель Vosk найдена!"; \
	else \
			echo "Модель Vosk не найдена.."; \
	fi