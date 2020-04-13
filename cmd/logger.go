package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"time"
)

type defaultFieldFormatter func(*logrus.Entry) ([]byte, error)

// Implements formatter interface
func (f defaultFieldFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Data["name"] = "Markr ✅"
	return f(e)
}

func initLogger() {
	formatter := log.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	}

	log.SetFormatter(defaultFieldFormatter(formatter.Format))
}
