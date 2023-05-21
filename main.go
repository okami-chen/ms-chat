package main

import (
	"context"
	"github.com/XM-GO/PandaKit/logger"
	"github.com/okamin-chen/service/pkg/config"
	"github.com/okamin-chen/service/pkg/global"
	"github.com/okamin-chen/service/pkg/initialize"
	"github.com/okamin-chen/service/pkg/starter"
	"github.com/okamin-chen/service/pkg/transport"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile string
	appEnv     string
)

var rootCmd = &cobra.Command{
	Use:   "chat is the main component in the chat.",
	Short: `chat is go gin frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if configFile == "" {
			global.Log.Panic("配置文件不存在")
		}
		global.Conf = config.InitConfig(configFile)
		global.Log = logger.InitLog(global.Conf.Log.File.GetFileName(), global.Conf.Log.Level)
		global.Db = starter.NewGorm()
		initialize.InitTable()
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
	rootCmd.Flags().StringVar(&appEnv, "env", getEnvStr("APP_ENV", ""), "app env.")
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
