package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
)

func NewViper() *viper.Viper {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	config := viper.New()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	return config
}
