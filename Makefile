run:
	cd cmd; go run main.go
max:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.42.zip
	unzip vosk-model-small-ru-0.42.zip
	mv vosk-model-ru-0.42 model_ru

small:
	wget https://alphacephei.com/vosk/models/vosk-model-small-ru-0.22.zip
	unzip vosk-model-small-ru-0.22.zip
	mv vosk-model-ru-0.22 model_ru

init: