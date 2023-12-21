package logic

import (
	"context"
	"errors"
	"github.com/Allen9012/Infinite/application/applet/internal/code"
	"github.com/Allen9012/Infinite/application/applet/internal/svc"
	"github.com/Allen9012/Infinite/application/applet/internal/types"
	"github.com/Allen9012/Infinite/application/user/rpc/user"
	"github.com/Allen9012/Infinite/pkg/encrypt"
	"github.com/Allen9012/Infinite/pkg/jwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixActivation = "biz#activation#%s"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 对注册请求中的信息预处理
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.RegisterPasswdEmpty
	} else {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}
	encedmobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: encedmobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, code.MobileHasRegistered
	}
	// 执行注册
	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Mobile:   encedmobile,
		Username: req.Name,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}
	// jwt存储session
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
		},
	})
	// 当生成了token，也就是已经验证好之后，删除缓存的验证码
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)
	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}

func checkVerificationCode(redis *redis.Redis, mobile string, code string) error {
	// 先获取验证码信息
	cacheCode, err := getActivationCache(mobile, redis)
	if err != nil {
		return err
	}
	// 注册但是发现验证码已经过期
	if code == "" {
		return errors.New("verification code expired")
	}
	// 验证码输入不正确
	if code != cacheCode {
		return errors.New("verification code expired")
	}
	return nil
}
