package hfctx

import (
	"context"
	"time"
)

type Ctx = context.Context

func DefaultCtx(ctx Ctx, defaultCtx Ctx) Ctx {
	if ctx == nil {
		return defaultCtx
	}
	return ctx
}

func Background() Ctx {
	return context.Background()
}

func TODO() Ctx {
	return context.TODO()
}

func WithValue(parent Ctx, key, val interface{}) Ctx {
	return context.WithValue(parent, key, val)
}

func WithValueDefault(parent Ctx, defaultCtx Ctx, key, val interface{}) Ctx {
	if parent == nil {
		return context.WithValue(defaultCtx, key, val)
	}
	return context.WithValue(parent, key, val)
}

func WithCancel(parent Ctx) (Ctx, context.CancelFunc) {
	return context.WithCancel(parent)
}

func WithCancelDefault(parent Ctx, defaultCtx Ctx) (Ctx, context.CancelFunc) {
	if parent == nil {
		return context.WithCancel(defaultCtx)
	}
	return context.WithCancel(parent)
}

func WithTimeout(parent Ctx, timeout time.Duration) (Ctx, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

func WithTimeoutDefault(parent Ctx, defaultCtx Ctx, timeout time.Duration) (Ctx, context.CancelFunc) {
	if parent == nil {
		return context.WithTimeout(defaultCtx, timeout)
	}
	return context.WithTimeout(parent, timeout)
}

func WithDeadline(parent Ctx, deadline time.Time) (Ctx, context.CancelFunc) {
	return context.WithDeadline(parent, deadline)
}

func WithDeadlineDefault(parent Ctx, defaultCtx Ctx, deadline time.Time) (Ctx, context.CancelFunc) {
	if parent == nil {
		return context.WithDeadline(defaultCtx, deadline)
	}
	return context.WithDeadline(parent, deadline)
}
