package logger

import (
	log "github.com/sirupsen/logrus"
)

func Warningf(format string, args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Warningf(format, args...)
}

func Warning(args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Warningln(args...)
}
