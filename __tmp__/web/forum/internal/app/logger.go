package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(viper.GetInt32("LOG.LEVEL")))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

type LogrusWriter struct {
	Logger *logrus.Logger
}

func (l *LogrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
