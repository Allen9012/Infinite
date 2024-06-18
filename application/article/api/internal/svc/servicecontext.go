package svc

import (
	"github.com/Allen9012/Infinite/application/article/api/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	defaultOssConnectTimeout   = 1
	defaultOssReadWriteTimeout = 3
)

type ServiceContext struct {
	Config    config.Config
	OssClient *oss.Client
	//ArticleRPC article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Oss.ConnectTimeout == 0 {
		c.Oss.ConnectTimeout = defaultOssConnectTimeout
	}
	if c.Oss.ReadWriteTimeout == 0 {
		c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	}
	// 初始化客户端
	oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret, oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		OssClient: oc,
		//ArticleRPC:
	}
}
