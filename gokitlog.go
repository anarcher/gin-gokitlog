package gokitlog

import (
	"time"

	"github.com/gin-gonic/gin"
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

// LogAdapter returns a new LogAdapter wrapper around the passed
// logger. It formats the log entries similar to
// http://godoc.org/github.com/gin-gonic/gin#Logger does.
//
// Example:
//        router := gin.New()
//        router.Use(gokitlog.LogAdapter(func(c *gin.Context){
//				 latency,_ := c.Get("latency")
//               logger.Log(
//					"latency",latency,"clientIP",c.ClientIP(),
//					"status",c.Writer.Status(),"err",c.Errors.String())
//		  })
func LogAdapter(fn func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		c.Set("latency", latency)
		fn(c)
	}
}
