package httpserver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
	_defaultAddr         = ":3120"
)

type Server struct {
	server *fiber.App
	port   string
}

func FiberConfig(status, appName string) fiber.Config {
	config := fiber.Config{
		AppName:               appName,
		DisableStartupMessage: true,
		EnablePrintRoutes:     false,
		ReadTimeout:           _defaultReadTimeout,
		WriteTimeout:          _defaultWriteTimeout,
		IdleTimeout:           60 * time.Second,
	}

	if status == "dev" {
		config.DisableStartupMessage = false
		config.EnablePrintRoutes = true
	}

	return config
}

func New(app *fiber.App, opts ...Option) *Server {
	s := &Server{
		server: app,
		port:   _defaultAddr,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}
func (s *Server) start() {
	idleConnsClosed := make(chan struct{})

	// TODO Graceful Shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		if err := s.server.Shutdown(); err != nil {
			fmt.Errorf("server shutdown error - %w", err)
		}

		close(idleConnsClosed)
	}()

	// Run server
	if err := s.server.Listen(s.port); err != nil {
		fmt.Errorf("server listen %w", err)
	}
}
