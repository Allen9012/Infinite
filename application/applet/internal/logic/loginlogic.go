package logic

import (
	"context"
	"github.com/Allen9012/Infinite/application/applet/internal/code"
	"github.com/Allen9012/Infinite/application/applet/internal/svc"
	"github.com/Allen9012/Infinite/application/applet/internal/types"
	"github.com/Allen9012/Infinite/application/user/rpc/user"
	"github.com/Allen9012/Infinite/pkg/encrypt"
	"github.com/Allen9012/Infinite/pkg/jwt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if req.VerificationCode == "" {
		return nil, code.VerificationCodeEmpty
	}
	// 手机验证码查找是否存在缓存中
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}
	encedmobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: encedmobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}

	// 构建 jwt的token
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
		},
	})
	if err != nil {
		return nil, err
	}
	// 成功之后删除激活码
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	return &types.LoginResponse{
		UserId: u.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
