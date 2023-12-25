package logic

import (
	"context"

	"github.com/Allen9012/Infinite/application/user/rpc/internal/svc"
	"github.com/Allen9012/Infinite/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *service.SendSmsRequest) (*service.SendSmsResponse, error) {
	// todo: send sms 使用aliyun的sdk

	return &service.SendSmsResponse{}, nil
}
