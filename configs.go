package main

import "github.com/spf13/viper"

func initiateConfigs() {
	initiateServerConfigs()
}

func initiateConfig(name string,directory string,configType string) *viper.Viper {
	cfg := viper.New()
	cfg.AddConfigPath(directory)
	cfg.SetConfigName(name)
	cfg.SetConfigType(configType)
	return cfg
}

func initiateServerConfigs() {
	cfg := initiateConfig("Server","./Configs","yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Configs = cfg
}
