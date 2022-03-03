package logger

import (
	log "github.com/sirupsen/logrus"
)

func Info(args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Infof(format, args...)
}
