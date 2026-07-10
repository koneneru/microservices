package logger

import (
	"log/slog"
	"os"

	"github.com/fatih/color"
)

type Options struct {
	Env         string
	ServiceName string
}

func New(opts *Options) *slog.Logger {
	var handler slog.Handler

	if opts.Env == "local" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       slog.LevelDebug,
			ReplaceAttr: tinitColors,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	logger := slog.New(handler).With(
		slog.String("source", opts.ServiceName),
		slog.String("env", opts.Env),
	)

	return logger
}

func tinitColors(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		switch level {
		case slog.LevelDebug:
			a.Value = slog.StringValue(color.New(color.FgBlue).Sprint("DBG"))
		case slog.LevelInfo:
			a.Value = slog.StringValue(color.New(color.FgGreen).Sprint("INF"))
		case slog.LevelWarn:
			a.Value = slog.StringValue(color.New(color.FgYellow).Sprint("WRN"))
		case slog.LevelError:
			a.Value = slog.StringValue(color.New(color.FgRed).Sprint("ERR"))
		}
	}

	if a.Key == slog.TimeKey {
		a.Value = slog.StringValue(a.Value.Time().Format("2006-01-02 15:04:05"))
	}
	return a
}
