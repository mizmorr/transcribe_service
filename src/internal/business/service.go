package business

import (
	"net/http"
	"os"
	"strings"

	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/transcriber/internal/domain"
)

type STT interface {
	Transcribe(filePath string, resultChan chan<- string, errChan chan<- string)
}

type Service struct {
	log *logger.Logger
	stt STT
}

func New(stt STT, log *logger.Logger) *Service {
	return &Service{
		stt: stt,
		log: log,
	}
}

func (svc *Service) Process(filePath string, ch chan<- *domain.Response) {

	chSuccess, chError := make(chan string), make(chan string)

	go svc.stt.Transcribe(filePath, chSuccess, chError)

	select {
	case success := <-chSuccess:
		ch <- &domain.Response{
			Status: http.StatusOK,
			Message: domain.Transcription{
				Text: success,
			},
		}
	case err := <-chError:
		ch <- &domain.Response{
			Status: http.StatusInternalServerError,
			Message: domain.Transcription{
				Text: "Some error occurs: " + err,
			},
		}
	}

}

func (svc *Service) CleanUp(filePath string) {
	if err := os.Remove(filePath); err != nil {
		svc.log.Err(err).Msg("Got error while removing audiofile " + filePath)
	}
	if strings.Contains(filePath, ".mp3") {

		wavFilePath := strings.Replace(filePath, ".mp3", ".wav", 1)

		if _, err := os.Stat(wavFilePath); err == nil {
			err = os.Remove(wavFilePath)
			if err != nil {
				svc.log.Err(err).Msg("Got error while removing audiofile " + wavFilePath)
			}
		}

	}

}
