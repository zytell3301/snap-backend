package main

import (
	"fmt"
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

	server.Get("/", func(ctx iris.Context) {
		fmt.Println("Request received")

		ctx.Write([]byte("Response returned"))
	})

	server.Listen(":"+Configs.GetString("port"),iris.WithLogLevel(Configs.GetString("log_level")))
}
