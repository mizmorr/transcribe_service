import sys
from vosk import Model, KaldiRecognizer
import wave

import logging

import json

import os

audio_file = sys.argv[1]

from vosk import SetLogLevel

SetLogLevel(-1)

if not os.path.exists("../../model_ru"):
    print("Ошибка: Модель 'model_ru' не найдена.")
    sys.exit(1)

if not audio_file.endswith(".wav"):
    print("Ошибка: Файл должен быть в формате WAV.")
    sys.exit(1)

try:
    wf = wave.open(audio_file, "rb")
except Exception as e:
    print(f"Ошибка при открытии файла: {e}")
    sys.exit(1)

sample_rate = wf.getframerate()

if sample_rate != 16000 and sample_rate != 8000:  
    print(f"Ошибка: Частота дискретизации {sample_rate} не поддерживается.")
    sys.exit(1)

model = Model("../../model_ru")

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