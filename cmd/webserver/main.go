package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/art-injener/otus-homework/internal/db/mysql"

	"github.com/gorilla/sessions"

	"github.com/art-injener/otus-homework/internal/service"

	"golang.org/x/sync/errgroup"

	"github.com/art-injener/otus-homework/internal/config"
	"github.com/art-injener/otus-homework/internal/logger"
	"github.com/art-injener/otus-homework/internal/repository/mysql/accounts"
	"github.com/art-injener/otus-homework/internal/rest"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := mysql.NewConnection(cfg)
	if err != nil {
		log.Println(err.Error())

		return
	}
	defer db.Close()

	repo := accounts.NewAccountsRepo(db)
	serv := service.NewUserService(repo, cfg.Log)
	session := sessions.NewCookieStore([]byte(cfg.SessionKey))
	webServer, err := rest.NewWebServer(serv, session, cfg)
	if err != nil {
		log.Println(err.Error())

		return
	}

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		err := webServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cfg.Log.Error().Msg("Start REST-API-server failed")
			_ = webServer.Shutdown(mainCtx)

			return err
		}

		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()

		return webServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %serv \n", err)
	}
}
