package main

import (
	"NewApiProd/internal/author"
	p1 "NewApiProd/internal/author/db/postgresql"
	"NewApiProd/internal/author/service"
	"NewApiProd/internal/config"
	"NewApiProd/internal/user"
	"NewApiProd/pkg/client/postgresql"
	"NewApiProd/pkg/logging"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%", err)
	}

	repository := p1.NewRepository(postgreSQLClient, logger)

	// bookRepository := db.NewRepository(postgreSQLClient, logger)

	// all, err := bookRepository.FindAll(context.TODO())
	// if err != nil {
	// 	logger.Fatalf("%", err)
	// }

	// for _, b := range all {
	// 	logger.Debug(b.Name)
	// }

	logger.Info("register author handler")
	authorService := service.NewService(repository, logger)
	authorHandler := author.NewHandler(authorService, logger)

	authorHandler.Register(router)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)

	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
