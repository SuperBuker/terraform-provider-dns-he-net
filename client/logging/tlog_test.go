package logging_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
)

func TestTLog(t *testing.T) {
	logger := logging.NewTlog()
	logger.Debug(t.Context(), "debug")
	logger.Error(t.Context(), "error")
	logger.Info(t.Context(), "info")
	logger.Trace(t.Context(), "trace")
	logger.Warn(t.Context(), "warn")
}
