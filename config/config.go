package config

import (
	"github.com/joho/godotenv"
	"os"
)

var Appconf = new(appConfig)

type appConfig struct {
	// 名称
	AppName string
	// 端口号
	AppPort string
	// 模式
	// debug & release
	AppMode string
	// 静态目录(本地目录)
	StaticPath string
}

func init() {
	godotenv.Load()
	Appconf.AppPort = os.Getenv("APP_PORT")
	Appconf.AppName = os.Getenv("APP_NAME")
	Appconf.StaticPath = os.Getenv("STATIC_PATH")
	Appconf.AppMode = os.Getenv("APP_MODE")
}
