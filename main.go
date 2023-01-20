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

//	@securityDefinitions.basic	BasicAuth
func main() {
	// Configuration
	var configFlag = flag.String("config", "./config/config.toml", "TOML file for configuration")
	/* [코드리뷰]
	 * 시스템과 관련된 config를 main 함수 첫번째에서 잘 가져와 주셨습니다.
	 * command 라인에서 config를 별도로 지정하여 환경에 따라 다른 시스템을 다르게 실행할 수 있게된 좋은 코드입니다.
	 */
	flag.Parse()
	cf := config.GetConfig(*configFlag)

	// Logger
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("InitLogger failed, err:%v\n", err)
		return
	}

	// Database
	db.ConnectDB(cf)

	// Server: start
	logger.Debug("Ready server")

	mapi := &http.Server{
		Addr:           cf.Server.Port,
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
