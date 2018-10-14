package trace

import (
	"context"

	"github.com/loov/unpolluted/analyser"
	"github.com/loov/unpolluted/trace/dtrace"
)

func Supported() error { return dtrace.Supported() }

func Program(ctx context.Context, analyser analyser.Analyser, cmd string, args ...string) (int, error) {
	return dtrace.Program(ctx, analyser, cmd, args)
}