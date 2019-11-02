package http

import (
	"context"
	"fmt"
	"log"
	_http "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crazy-genius/imgresizer/internal/imgresizer/configuration"
	"github.com/gin-gonic/gin"
)

// Server represents resize  httpserver instance
type Server struct {
	server _http.Server
}

// NewServer create new Server
func NewServer(c configuration.Configuration) *Server {
	router := gin.Default()
	if c.EnableRateLimit {
		router.Use(limitConnectionsMiddleWare(c.MaxConnections))
	}

	router.Use(gin.Recovery())
	router.Use(versionMiddleWare(configuration.VERSION))
	router.Use(requestIDMiddleWare())

	router.GET("/", hello)
	router.GET("/resize", resize)

	return &Server{
		server: _http.Server{
			Addr:    fmt.Sprintf("%s:%d", c.Host, c.Port),
			Handler: router,
		},
	}
}

// StartAndListenSignals start listening on desired port
func (s *Server) StartAndListenSignals() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != _http.ErrBodyReadAfterClose {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
