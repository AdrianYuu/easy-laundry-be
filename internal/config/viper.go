package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
)

func NewViper() *viper.Viper {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Errorf("failed to load env: %v", err))
	}

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	return v
}
