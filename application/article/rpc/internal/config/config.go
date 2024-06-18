package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string // 数据库的dsn
	//CacheRedis cache.CacheConf
	//BizRedis   redis.RedisConf
	//Consul     consul.Conf
}
