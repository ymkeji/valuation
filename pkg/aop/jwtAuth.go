package aop

import (
	"context"

	"valuation/pkg/errorx"
	"valuation/pkg/jwt"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func JWTAuthMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				jwtToken := header.RequestHeader().Get("Authorization")
				id, err := jwt.DecodeToken(jwtToken)
				errorx.Dangerous(err)
				ctx = context.WithValue(ctx, "id", id)
				return handler(ctx, req)
			}
			return nil, err
		}
	}
}
