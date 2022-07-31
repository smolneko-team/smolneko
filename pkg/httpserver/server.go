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
    _defaultAddr         = ":80"
)

type Server struct {
    server *fiber.App
    port   string
}

func FiberConfig() fiber.Config {
    return fiber.Config{
        ReadTimeout:  _defaultReadTimeout,
        WriteTimeout: _defaultWriteTimeout,
    }
}

func New(app *fiber.App, opts ...Option) *Server {

    s := &Server{
        server: app,
        port:   _defaultAddr,x
    }

    for _, opt := range opts {
        opt(s)
    }

    // TODO start simple server without graceful shutdown for dev purposes (stage status env)

    s.start()

    return s
}
func (s *Server) start() {
    idleConnsClosed := make(chan struct{})

    // TODO Graceful Shutdown ?
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
    // TODO TLS ?
    if err := s.server.Listen(s.port); err != nil {
        fmt.Errorf("server listen %w", err)
    }
}
