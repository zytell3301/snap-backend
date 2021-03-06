package main

import (
	"github.com/spf13/viper"
	"snap/Database/Cassandra"
	"snap/Database/Redis"
	"snap/Database/Uuid"
	"snap/DriversService"
)

/**
 * TODO: Add an example file of every config with comments named like CONFIG_NAME.CONFIG_EXTENSION.example. e.g Redis.yaml.example OR Server.yaml.example
 */

func initiateConfigs() {
	initiateServerConfigs()
	initiateRedisConfigs()
	initiateGrpcConfigs()
	initiateUuidConfigs()
	initiateCassandraConfigs()
}

func initiateConfig(name string, directory string, configType string) *viper.Viper {
	cfg := viper.New()
	cfg.AddConfigPath(directory)
	cfg.SetConfigName(name)
	cfg.SetConfigType(configType)
	return cfg
}

func initiateServerConfigs() {
	cfg := initiateConfig("Server", "./Configs", "yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Configs = cfg
}

func initiateRedisConfigs() {
	cfg := initiateConfig("Redis", "./Configs", "yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Redis.Configs = cfg
}

func initiateGrpcConfigs() {
	cfg := initiateConfig("Grpc", "./Configs", "yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	DriversService.Configs = cfg
}

func initiateUuidConfigs() {
	cfg := initiateConfig("Uuid", "./Configs", "yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Uuid.Configs = cfg
}

func initiateCassandraConfigs() {
	cfg := initiateConfig("Cassandra", "./Configs", "yaml")
	err := cfg.ReadInConfig()

	switch err != nil {
	case true:
		panic(err)
	}

	Cassandra.Configs = cfg
}
