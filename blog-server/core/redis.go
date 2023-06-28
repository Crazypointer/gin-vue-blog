package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"server/global"
	"strconv"
	"time"
)

// ConnectRedis 默认连接0号数据库
func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

// ConnectRedisDB 指定连接的数据库
func ConnectRedisDB(db int) *redis.Client {
	addr := global.Config.Redis.IP + ":" + strconv.Itoa(global.Config.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.Config.Redis.Password, // no password set
		DB:       0,                            // use default DB
		PoolSize: global.Config.Redis.PoolSize, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis 链接失败: %s", addr)
	}
	return rdb
}
