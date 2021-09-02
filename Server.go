package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

var Configs *viper.Viper

func init() {
	initiateConfigs()
	initiatePackages()
}

func main() {
	server := iris.New()

	server.Listen(":"+Configs.GetString("port"),iris.WithLogLevel(Configs.GetString("log_level")))
}
