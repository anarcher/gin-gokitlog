package gokitlog

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
)

// ErrorLogger returns an ErrorLoggerT with parameter gin.ErrorTypeAny
func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

// ErrorLoggerT returns an ErrorLoggerT middleware with the given
// type gin.ErrorType.
func ErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() {
			json := c.Errors.ByType(typ).JSON()
			if json != nil {
				c.JSON(-1, json)
			}
		}
	}
}

// Logger prints a logline for each request and measures the time to
// process for a call. It formats the log entries similar to
// http://godoc.org/github.com/gin-gonic/gin#Logger does.
//
// Example:
//        router := gin.New()
//        router.Use(ginglog.Logger(logger))
func Logger(plogger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.NewContext(plogger)
		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		logger.Log("latency", latency, "clientIP", clientIP, "status", statusCode, "method", method, "path", path, "err", c.Errors.String())

	}
}
