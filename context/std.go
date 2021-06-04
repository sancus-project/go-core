package context

import (
	"context"
)

type (
	Context = context.Context
)

func Background() context.Context {
	return context.Background()
}

func WithValue(ctx context.Context, key interface{}, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

