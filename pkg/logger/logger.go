package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var once sync.Once

// GetLogger возвращает синглтон логгера
func GetLogger() *logrus.Logger {
    once.Do(func() {
        log = logrus.New()
        log.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
        })
        log.SetLevel(logrus.InfoLevel)
    })
    return log
}
