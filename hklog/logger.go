package hklog

import "go.hikit.io/hikit/hkctx"

type Logger interface {
	Info(ctx hkctx.Ctx)
}
