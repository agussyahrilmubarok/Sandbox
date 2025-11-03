package app

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenPostgres(cfg *viper.Viper, log *logrus.Logger) *gorm.DB {

	host := cfg.GetString("POSTGRES.HOST")
	port := cfg.GetInt("POSTGRES.PORT")
	user := cfg.GetString("POSTGRES.USER")
	password := cfg.GetString("POSTGRES.PASSWORD")
	name := cfg.GetString("POSTGRES.NAME")
	sslMode := cfg.GetString("POSTGRES.SSL_MODE")
	maxOpenConns := cfg.GetInt("POSTGRES.MAX_OPEN_CONNS")
	maxIdleConns := cfg.GetInt("POSTGRES.MAX_IDLE_CONNS")
	maxIdleTime := cfg.GetDuration("POSTGRES.MAX_IDLE_TIME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user, password, host, port, name, sslMode)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&LogrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := database.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(maxIdleConns)
	connection.SetMaxOpenConns(maxOpenConns)
	connection.SetConnMaxLifetime(time.Second * maxIdleTime)

	return database
}
