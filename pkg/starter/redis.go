package starter

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/okamin-chen/chat/pkg/global"
	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.Host + global.Conf.Redis.GetPort(),
		Password: global.Conf.Redis.Password, // no password set
		DB:       global.Conf.Redis.Db,       // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Log.Panic("Redis Init Fail")
		return nil
	}
	global.Log.Infoln("Redis Init Success")
	return rdb
}

func NewLock() *redsync.Redsync {
	pool := goredis.NewPool(global.Redis)
	rc := redsync.New(pool)
	global.Log.Infoln("Lock Init Success")
	return rc
}
