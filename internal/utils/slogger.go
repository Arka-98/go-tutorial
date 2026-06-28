package utils

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type ColoredHandler struct {
	slog.Handler
}

func (c *ColoredHandler) Handle(ctx context.Context, r slog.Record) error {
	var color string

	switch r.Level {
	case slog.LevelDebug:
		color = blue
	case slog.LevelInfo:
		color = green
	case slog.LevelWarn:
		color = yellow
	case slog.LevelError:
		color = red
	}

	os.Stdout.Write([]byte(color))
	defer os.Stdout.Write([]byte(reset))

	return c.Handler.Handle(ctx, r)
}

func (h *ColoredHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ColoredHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *ColoredHandler) WithGroup(name string) slog.Handler {
	return &ColoredHandler{Handler: h.Handler.WithGroup(name)}
}

func SloggerEx(out io.Writer) error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, ok := a.Value.Any().(*slog.Source)

				if !ok {
					return a
				}

				if rel, err := filepath.Rel(cwd, source.File); err == nil {
					source.File = rel
				}
			}

			return a
		},
	}
	ch := &ColoredHandler{Handler: slog.NewTextHandler(os.Stdout, opts)}
	logger := slog.New(slog.NewMultiHandler(ch, slog.NewJSONHandler(out, opts)))

	slog.SetDefault(logger)

	return nil
}
