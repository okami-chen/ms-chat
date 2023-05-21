package starter

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/okamin-chen/service/pkg/global"
)

func NewWechatMiniProgram() *miniProgram.MiniProgram {
	MiniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     global.Conf.Wechat.MiniProgram.AppId,     // 小程序appid
		Secret:    global.Conf.Wechat.MiniProgram.AppSecret, // 小程序app secret
		HttpDebug: global.Conf.Wechat.MiniProgram.HttpDebug,
		Debug:     global.Conf.Wechat.MiniProgram.Debug,
		AESKey:    global.Conf.Wechat.MiniProgram.AesKey,
		Log: miniProgram.Log{
			Level: global.Conf.Wechat.MiniProgram.Level,
			File:  global.Conf.Wechat.MiniProgram.File,
		},
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     global.Conf.Cache.Host + global.Conf.Cache.GetPort(),
			Password: global.Conf.Cache.Password,
			DB:       global.Conf.Cache.Db,
		}),
	})
	if err != nil {
		global.Log.Panic(err)
	}
	global.Log.Infoln("Wechat MiniProgram Init Success")
	return MiniProgramApp
}
