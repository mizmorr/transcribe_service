package app

import (
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/transcriber/internal/business"
	"github.com/mizmorr/transcriber/internal/delivery"
	"github.com/mizmorr/transcriber/internal/transcriber"
	"github.com/mizmorr/transcriber/pkg/server"
)

type lifecycle interface {
	Start(context.Context) error
	Stop(context.Context) error
}
type App struct {
	log     *logger.Logger
	workers []lifecycle
}

func New() *App {
	return &App{
		log: logger.Get("logs/app.log", "debug"),
	}
}

func (a *App) Start(ctx context.Context) error {

	if _, ok := ctx.Value("logger").(*logger.Logger); !ok {
		ctx = context.WithValue(ctx, "logger", a.log)
	}

	stt := new(transcriber.Transcriber)

	service := business.New(stt, a.log)

	controller := delivery.NewController(service)

	handler := gin.New()

	delivery.NewRouter(handler, controller)

	server := server.New(handler, "0.0.0.0", "8080", time.Second*10)

	if err := server.Start(ctx); err != nil {
		return err
	}
	a.workers = append(a.workers, server)

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	if _, ok := ctx.Value("logger").(*logger.Logger); !ok {
		ctx = context.WithValue(ctx, "logger", a.log)
	}

	chanDone := make(chan any)
	var wg sync.WaitGroup

	for _, worker := range a.workers {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			chanDone <- worker.Stop(ctx)

		}(&wg)
	}

	go func() {
		wg.Wait()
		close(chanDone)
	}()

	for res := range chanDone {
		if res != nil {
			a.log.Debug().Msg(res.(string))
		}
	}
	return nil
}
