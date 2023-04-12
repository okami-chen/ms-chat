package starter

import (
	"context"
	"github.com/okamin-chen/chat/pkg/global"
	"github.com/redis/go-redis/v9"
)

func NewCache() *redis.Client {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Conf.Cache.Host + global.Conf.Cache.GetPort(),
		Password: global.Conf.Cache.Password, // no password set
		DB:       global.Conf.Cache.Db,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Log.Warning(err.Error())
		global.Log.Panic("Redis Init Fail")
		return nil
	}
	global.Log.Infoln("Cache Init Success")
	return rdb
}
