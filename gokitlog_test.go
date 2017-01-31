package gokitlog

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

func Test_Logger(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	middleware := Logger(logger)

	if middleware == nil {
		t.Errorf("Can't get Logger middleware")
	}
}
