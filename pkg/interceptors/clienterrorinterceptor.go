package interceptors

import (
	"context"
	"github.com/Allen9012/Infinite/pkg/xcode"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ClientErrorInterceptor 获取grpc错误，detail中解析出自定义的业务错误，拿到错误后转化为xcode
func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			grpcStatus, _ := status.FromError(err)
			xc := xcode.GrpcStatusToXCode(grpcStatus)
			err = errors.WithMessage(xc, grpcStatus.Message())
		}

		return err
	}
}
