package Uuid

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var space uuid.UUID
var Configs *viper.Viper

func InitSpace() {
	space = uuid.New()
	err := space.UnmarshalText([]byte(Configs.GetString("space")))

	switch err != nil {
	case true:
		panic(err)
	}
}

func GenerateV5(name string) string {
	return uuid.NewSHA1(space, []byte(name)).String()
}
