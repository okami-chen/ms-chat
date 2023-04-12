package main

import (
	"context"
	"github.com/XM-GO/PandaKit/logger"
	"github.com/okamin-chen/chat/pkg/config"
	"github.com/okamin-chen/chat/pkg/global"
	"github.com/okamin-chen/chat/pkg/initialize"
	"github.com/okamin-chen/chat/pkg/starter"
	"github.com/okamin-chen/chat/pkg/transport"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "chat is the main component in the chat.",
	Short: `chat is go gin frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if configFile != "" {
			global.Conf = config.InitConfig(configFile)
			global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
			dbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
			dbGorm.Dsn = global.Conf.Mysql.Dsn()
			dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
			dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
			global.Db = dbGorm.GormInit()
			initialize.InitTable()
		} else {
			global.Log.Panic("请配置config")
		}
		global.Cache = starter.NewCache()
		global.Redis = starter.NewRedis()
		global.Lock = starter.NewLock()
		global.MiniProgram = starter.NewWechatMiniProgram()
	},
	Run: func(cmd *cobra.Command, args []string) {
		app := transport.NewHttpServer()
		app.Start(context.TODO())
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop

		if err := app.Stop(context.TODO()); err != nil {
			log.Fatalf("fatal app stop: %s", err)
			os.Exit(-3)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", getEnvStr("CONFIG", "./config.yml"), "config file path.")
}

func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("panda root cmd execute: %s", err)
		os.Exit(1)
	}
}
