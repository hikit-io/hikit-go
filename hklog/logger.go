package hklog

import "go.hikit.io/hkctx"

type Logger interface {
	Info(ctx hkctx.Ctx)
}
