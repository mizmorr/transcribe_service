package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/mizmorr/transcriber/internal/app"
)

func main() {
	if err := execute(); err != nil {
		panic(err)
	}

}

func execute() error {

	go func() {
		fmt.Println("Starting pprof on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof error:", err)
		}
	}()

	ctx := context.Background()

	app := app.New()

	if err := app.Start(ctx); err != nil {
		return err
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-stopCh

	return app.Stop(ctx)
}
