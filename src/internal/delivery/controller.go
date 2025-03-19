package delivery

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/transcriber/internal/domain"
)

type Service interface {
	Process(string, chan<- *domain.Response)
	CleanUp(string)
}

type Controller struct {
	svc Service
}

func NewController(svc Service) *Controller {
	return &Controller{
		svc: svc,
	}
}
func (c *Controller) Transcribe(g *gin.Context) {
	file, err := g.FormFile("audio")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Файл не найден"})
		return
	}
	savePath := filepath.Join("../internal/store/", file.Filename)

	if err := g.SaveUploadedFile(file, savePath); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	defer c.svc.CleanUp(savePath)

	resultCh := make(chan *domain.Response)

	go c.svc.Process(savePath, resultCh)

	result := <-resultCh

	g.JSON(result.Status, result.Message)
}
