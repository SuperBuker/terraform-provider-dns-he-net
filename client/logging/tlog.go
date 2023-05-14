package logging

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewTlog() Logger {
	return &tlogLogger{}
}

type tlogLogger struct{}

func (tlogLogger) Debug(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Debug(ctx, msg, additionalFields...)
}

func (tlogLogger) Error(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Error(ctx, msg, additionalFields...)
}

func (tlogLogger) Fatal(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Error(ctx, msg, additionalFields...)
}

func (tlogLogger) Info(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Info(ctx, msg, additionalFields...)
}

func (tlogLogger) Panic(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Error(ctx, msg, additionalFields...)
}

func (tlogLogger) Trace(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Trace(ctx, msg, additionalFields...)
}

func (tlogLogger) Warn(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tflog.Warn(ctx, msg, additionalFields...)
}
