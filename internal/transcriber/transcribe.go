package transcriber

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/mizmorr/transcriber/pkg/utils"
)

type Transcriber struct{}

func (t *Transcriber) Transcribe(filePath string, resultChan chan<- string, errChan chan<- string) {

	if err := t.preProcessing(filePath); err != nil {
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

func (t *Transcriber) preProcessing(filePath string) error {
	isMp3, err := utils.IsMP3(filePath)

	if err != nil {
		return err
	}

	if isMp3 {
		err = utils.ConvertMP3ToWAV(filePath, strings.Replace(filePath, ".mp3", ".wav", 1))

		if err != nil {
			return err
		}
	}

	isWav, err := utils.IsWAV(filePath)

	if err != nil {
		return err

	}

	if !isWav {
		return errors.New("unexpected audio format")
	}

	return nil
}
