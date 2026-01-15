package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogrus(viper *viper.Viper) *logrus.Logger {
	l := logrus.New()

	l.SetLevel(logrus.Level(viper.GetInt("log.level")))
	l.SetFormatter(&logrus.JSONFormatter{})

	return l
}
