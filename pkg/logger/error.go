package logger

import (
	log "github.com/sirupsen/logrus"
)

func Error(args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Errorf(format, args...)
}
