package logging_test

import (
	"errors"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/rs/zerolog"
)

func TestZeroLog(t *testing.T) {
	logger := logging.NewZerolog(zerolog.TraceLevel, false)

	fields := logging.Fields{
		"string":    "string",
		"int":       1,
		"int8":      int8(1),
		"int16":     int16(1),
		"int32":     int32(1),
		"int64":     int64(1),
		"uint":      uint(1),
		"uint8":     uint8(1),
		"uint16":    uint16(1),
		"uint32":    uint32(1),
		"uint64":    uint64(1),
		"float32":   float32(1),
		"float64":   float64(1),
		"bool":      true,
		"error":     errors.New("error"),
		"interface": nil,
	}

	logger.Debug(t.Context(), "debug", fields)
	logger.Error(t.Context(), "error", fields)
	logger.Info(t.Context(), "info", fields)
	logger.Trace(t.Context(), "trace", fields)
	logger.Warn(t.Context(), "warn", fields)

	logger.Debug(t.Context(), "", fields)
}
