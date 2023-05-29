// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/evrone/go-clean-template/internal"
	"os"
	"os/signal"
	"syscall"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/httpserver"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	log := internal.InitializeLogger()

	//servers
	rmqServer := internal.InitializeNewRmqRpcServerWithConfig(cfg)
	httpServer, _ := httpserver.New(cfg)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		log.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
