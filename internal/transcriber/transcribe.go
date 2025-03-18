package transcriber

import (
	"os/exec"
)

type Transcriber struct{}

func (t *Transcriber) Transcribe(filePath string, resultChan chan<- string, errChan chan<- string) {

	cmd := exec.Command("python", "../internal/transcriber/transcribe.py", "../output.wav")

	output, err := cmd.CombinedOutput()
	if err != nil {
		errChan <- "error processing audio: " + err.Error()
		return
	}

	resultChan <- string(output)
}
