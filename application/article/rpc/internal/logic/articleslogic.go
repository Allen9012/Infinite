package logic

import (
	"context"

	"github.com/Allen9012/Infinite/application/article/rpc/internal/svc"
	"github.com/Allen9012/Infinite/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticlesLogic) Articles(in *pb.ArticlesRequest) (*pb.ArticlesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ArticlesResponse{}, nil
}
