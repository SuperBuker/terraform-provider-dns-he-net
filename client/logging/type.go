package logging

import "context"

// Logger is a generic interface for logging
type Logger interface {
	Debug(ctx context.Context, msg string, additionalFields ...map[string]interface{})
	Error(ctx context.Context, msg string, additionalFields ...map[string]interface{})
	Info(ctx context.Context, msg string, additionalFields ...map[string]interface{})
	Trace(ctx context.Context, msg string, additionalFields ...map[string]interface{})
	Warn(ctx context.Context, msg string, additionalFields ...map[string]interface{})
}

// Fields is a map of fields to be logged
type Fields map[string]interface{}
