package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/art-injener/otus/internal/config"
	"github.com/art-injener/otus/internal/logger"
	"github.com/art-injener/otus/internal/rest"
)

func main() {
	// читаем конфигурационные настройки
	cfg, err := config.LoadConfig("configs")
	if err != nil {
		log.Println(err.Error())

		return
	}

	cfg.Log = logger.NewConsole(cfg.LogLevel == config.DebugLevel)

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer, err := rest.CreateWebServer(cfg)
	if err != nil {
		log.Println(err.Error())

		return
	}

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cfg.Log.Error().Msg("Start REST-API-server failed")
			_ = httpServer.Shutdown(mainCtx)

			return err
		}

		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()

		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}

}
