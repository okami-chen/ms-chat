package global

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/go-redsync/redsync/v4"
	"github.com/okamin-chen/chat/pkg/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log         *logrus.Logger // 日志
	Db          *gorm.DB       // gorm
	Conf        *config.Config
	Cache       *redis.Client
	Redis       *redis.Client
	Lock        *redsync.Redsync
	MiniProgram *miniProgram.MiniProgram
)
