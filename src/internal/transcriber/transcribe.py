import sys
from vosk import Model, KaldiRecognizer
import wave

import logging

import json

import os

audio_file = sys.argv[1]

from vosk import SetLogLevel

SetLogLevel(-1)

wf = wave.open(audio_file, "rb")

sample_rate = wf.getframerate()

#docker conteiner layout
model_path = "./model_ru"

model = Model(model_path)

wf = wave.open(audio_file, "rb")

recognizer = KaldiRecognizer(model, wf.getframerate())

result_text = ""

while True:
    data = wf.readframes(4000)
    if len(data) == 0:
        break
    if recognizer.AcceptWaveform(data):
        result = json.loads(recognizer.Result())  
        result_text += result.get("text", "") + " "  

# Добавляем финальный результат
final_result = json.loads(recognizer.FinalResult())
result_text += final_result.get("text", "")

print(result_text.strip())

wf.close()