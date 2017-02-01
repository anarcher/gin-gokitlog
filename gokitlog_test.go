package gokitlog

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
)

func Test_Logger(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	middleware := LogAdapter(func(c *gin.Context) {
		latency, _ := c.Get("latency")
		logger.Log("a", "b", "latency", latency)
	})

	if middleware == nil {
		t.Errorf("Can't get Logger middleware")
	}
}
