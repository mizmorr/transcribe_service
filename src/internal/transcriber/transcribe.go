package transcriber

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/mizmorr/transcriber/pkg/utils"
)

type Transcriber struct{}

func (t *Transcriber) Transcribe(filePath string, resultChan chan<- string, errChan chan<- string) {
	var err error

	if filePath, err = t.preProcessing(filePath); err != nil {
		errChan <- err.Error()
		return
	}

	cmd := exec.Command("python", "../internal/transcriber/transcribe.py", filePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		errChan <- "error processing audio: " + err.Error()
		return
	}

	resultChan <- string(output)
}

func (t *Transcriber) preProcessing(filePath string) (string, error) {
	isWav, err := utils.IsWAV(filePath)

	if err != nil {
		return "", err

	}

	if isWav {
		return filePath, nil
	}

	isMp3, err := utils.IsMP3(filePath)

	if err != nil {
		return "", err

	}
	if isMp3 {
		wavPath := strings.Replace(filePath, ".mp3", ".wav", 1)
		return wavPath, utils.ConvertMP3ToWAV(filePath, wavPath)
	}

	return "", errors.New("unexpected audio format")
}
