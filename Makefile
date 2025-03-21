install-max-model:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.42.zip
	unzip vosk-model-small-ru-0.42.zip
	mv vosk-model-ru-0.42 model_ru

install-small-model:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.22.zip
	unzip vosk-model-small-ru-0.22.zip
	mv vosk-model-ru-0.22 model_ru

check:
	@if [ -d "model_ru" ]; then \
			echo "Модель Vosk найдена!"; \
	else \
			echo "Модель Vosk не найдена.."; \
	fi