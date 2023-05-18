package logging_test

import (
	"context"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
)

func TestTLog(t *testing.T) {
	logger := logging.NewTlog()
	logger.Debug(context.Background(), "debug")
	logger.Error(context.Background(), "error")
	logger.Info(context.Background(), "info")
	logger.Trace(context.Background(), "trace")
	logger.Warn(context.Background(), "warn")
}
