package main

import "github.com/spf13/viper"

func initiateConfigs() {
	initiateServerConfigs()
}

func initiateServerConfigs() {
	cfg := viper.New()
	cfg.AddConfigPath("./Configs")
	cfg.SetConfigName("Server")
	cfg.SetConfigType("yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Configs = cfg
}
