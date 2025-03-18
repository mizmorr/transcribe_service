package business

import (
	"net/http"
	"os"

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

}
