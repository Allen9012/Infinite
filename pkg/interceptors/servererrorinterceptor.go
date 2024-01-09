package interceptors

import (
	"context"
	"github.com/Allen9012/Infinite/pkg/xcode"
	"google.golang.org/grpc"
)

// ServerErrorInterceptor 通过拦截器来获得自定义错误类型
func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.FromError(err).Err()
	}
}
