package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var mp3Signatures = [][]byte{
	{0xFF, 0xFB}, // MPEG-1 Layer III
	{0xFF, 0xF3}, // MPEG-2 Layer III
	{0xFF, 0xF2}, // MPEG-2.5 Layer III
	{0xFF, 0xE3},
}

func IsMP3(filePath string) (bool, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return false, errors.New("Ошибка открытия файла:" + err.Error())

	}
	defer file.Close()

	header := make([]byte, 2)
	_, err = file.Read(header)
	if err != nil {
		return false, errors.New("Ошибка чтения файла:" + err.Error())

	}

	for _, sig := range mp3Signatures {
		if bytes.Equal(header, sig) {
			return true, nil
		}
	}

	return false, nil
}

type WAVHeader struct {
	RIFFHeader [4]byte // Должно быть "RIFF"
	FileSize   uint32  // Размер файла минус 8 байт
	WAVEHeader [4]byte // Должно быть "WAVE"
}

func IsWAV(filePath string) (bool, error) {
	file, _ := os.Open(filePath)

	defer file.Close()

	var header WAVHeader
	err := binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return false, fmt.Errorf("ошибка при чтении заголовка WAV: %v", err)
	}

	if !bytes.Equal(header.RIFFHeader[:], []byte("RIFF")) ||
		!bytes.Equal(header.WAVEHeader[:], []byte("WAVE")) {
		return false, nil
	}

	return true, nil
}

func ConvertMP3ToWAV(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ar", "16000", "-ac", "1", outputPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка при конвертации: %v — %s", err, stderr.String())
	}
	return nil
}
