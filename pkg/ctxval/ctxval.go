package ctxval

import (
	"context"

	"github.com/mcholismalik/boilerplate-golang/internal/model/abstraction"
)

type key string

var (
	keyAuth key = "x-auth"
	keyTrx  key = "x-trx"
)

func SetAuthValue(ctx context.Context, payload *abstraction.AuthContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetAuthValue(ctx context.Context) *abstraction.AuthContext {
	return ctx.Value(keyAuth).(*abstraction.AuthContext)
}

func SetTrxValue(ctx context.Context, payload *abstraction.TrxContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetTrxValue(ctx context.Context) *abstraction.TrxContext {
	return ctx.Value(keyTrx).(*abstraction.TrxContext)
}
