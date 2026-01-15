package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewGorm(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	dbUrl := viper.GetString("db.url")
	idleConnection := viper.GetInt("db.pool.idle")
	maxConnection := viper.GetInt("db.pool.max")
	lifetimeConnection := viper.GetInt("db.pool.lifetime")

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		})},
	)

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	conn, err := db.DB()

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	conn.SetMaxIdleConns(idleConnection)
	conn.SetMaxOpenConns(maxConnection)
	conn.SetConnMaxLifetime(time.Second * time.Duration(lifetimeConnection))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
