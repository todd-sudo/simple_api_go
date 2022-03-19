package logger

import (
	log "github.com/sirupsen/logrus"
)

func Fatal(args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	if !loggerDebug {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.Fatalf(format, args...)
}
