package cache

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/iter-x/iter-x/internal/conf"
)

type (
	Cacher interface {
		Close() error
		Client() *redis.Client
	}

	cache struct {
		cli *redis.Client
	}
)

func (c *cache) Close() error {
	return c.cli.Close()
}

func (c *cache) Client() *redis.Client {
	return c.cli
}

// NewCache new cache
func NewCache(c *conf.Data) Cacher {
	cacheConf := c.GetCache()
	redisConf := cacheConf.GetRedis()
	switch cacheConf.GetDriver() {
	case conf.Data_Cache_REDIS:
		return &cache{cli: newRedisCli(redisConf.GetAddr(), redisConf)}
	default:
		cli, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		return newRedisCacherByMiniRedis(cli, redisConf)
	}
}

// newRedisCacherByMiniRedis creates a new redis cacher by mini redis
func newRedisCacherByMiniRedis(cli *miniredis.Miniredis, redisConf *conf.Redis) Cacher {
	return &cache{cli: newRedisCli(cli.Addr(), redisConf)}
}

// newRedisCli creates a new redis client
func newRedisCli(addr string, redisConf *conf.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     redisConf.GetPassword(),
		DB:           int(redisConf.GetDb()),
		WriteTimeout: redisConf.GetWriteTimeout().AsDuration(),
		ReadTimeout:  redisConf.GetReadTimeout().AsDuration(),
		DialTimeout:  redisConf.GetDialTimeout().AsDuration(),
		Network:      redisConf.GetNetwork(),
	})
}
