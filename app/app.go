package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/gin-gonic/gin"
	"github.com/naipad/loach"
)

type Application struct {
	HttpServer *http.Server
	Router     *gin.Engine
	DB         *loach.DB
	MemCache   *fastcache.Cache
}

func New(addr, dbpath string) *Application {
	app := new(Application)
	app.Router = gin.New()
	app.Router.RedirectTrailingSlash = true

	app.HttpServer = &http.Server{
		Addr:    addr,
		Handler: app.Router.Handler(),
	}
	return app
}

func (app *Application) Start() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := app.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	log.Printf("Application running. Listening on %v\n", app.HttpServer.Addr)

	sig := <-sigs
	log.Printf("Received signal: %s. Shutting down...\n", sig)
	app.Close()

	wg.Wait()
}

func (app *Application) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.HttpServer.Shutdown(ctx); err != nil {
		log.Println("HTTP server Shutdown failed:", err)
	}

	if app.DB != nil {
		if err := app.DB.Close(); err != nil {
			log.Println("Database disconnect failed:", err)
		}
	}
}
