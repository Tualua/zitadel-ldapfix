package middleware

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Tualua/zitadel-ldapfix/internal/api/service"
	_ "github.com/Tualua/zitadel-ldapfix/internal/statik"
)

func ServiceHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		namer, ok := info.Server.(interface{ AppName() string })
		if !ok {
			return handler(ctx, req)
		}
		ctx = service.WithService(ctx, namer.AppName())
		return handler(ctx, req)
	}
}
