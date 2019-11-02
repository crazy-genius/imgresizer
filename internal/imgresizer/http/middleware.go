package http

import (
	_http "net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func versionMiddleWare(version string) gin.HandlerFunc {

	revision := strings.TrimSpace(version)

	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Revision", revision)
		c.Next()
	}
}

func requestIDMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func limitConnectionsMiddleWare(maxConnections int) gin.HandlerFunc {

	sem := make(chan struct{}, maxConnections)

	acqure := func(sem chan<- struct{}) {
		sem <- struct{}{}
	}

	release := func(sem <-chan struct{}) {
		<-sem
	}

	return func(c *gin.Context) {
		acqure(sem)
		defer release(sem)

		c.Next()
	}
}

func checkHostMiddleware(allowedHosts []string) gin.HandlerFunc {

	return func(c *gin.Context) {

		for _, allowedHost := range allowedHosts {
			if allowedHost == "*" {
				c.Next()
				return
			}
			if c.Request.URL.Hostname() == allowedHost {
				c.Next()
				return
			}
		}

		c.AbortWithStatus(_http.StatusUnauthorized)
	}
}
