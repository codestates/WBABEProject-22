package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"

	"oos/config"
	"oos/db"
	"oos/logger"
	"oos/router"
)

var g errgroup.Group

//	@title			Online ordering system
//	@version		1.0
//	@description	A simple online ordering system written in Go

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	// Configuration
	var configFlag = flag.String("config", "./config/config.toml", "TOML file for configuration")
	flag.Parse()
	cfg, err := config.GetConfig(*configFlag)
	if err != nil {
		fmt.Printf("GetConfig failed, err:%v\n", err)
		logger.Fatal("Error configuration file")
		return
	}

	// Logger
	if err := logger.InitLogger(cfg); err != nil {
		fmt.Printf("InitLogger failed, err:%v\n", err)
		logger.Fatal("Error loading logger")
		return
	}

	// Environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Loading .env failed, err:%v\n", err)
        logger.Fatal("Error loading .env file")
		return
    }

	// Database
	db.ConnectDB(cfg)

	// Server: start
	logger.Debug("Ready server")

	mapi := &http.Server{
		Addr:           cfg.Server.Port,
		Handler:        router.Engine(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		return mapi.ListenAndServe()
	})

	// Server: graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	logger.Warn("Shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mapi.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown:", err)
	}

	select {
	case <-ctx.Done():
		logger.Info("Timeout 5 seconds")
	default:
	}

	logger.Info("Server exiting")

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}
}
