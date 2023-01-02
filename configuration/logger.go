package configuration

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	file, err := os.OpenFile("logs/log_"+date.Format("01-02-2006_15")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicLogging(err)
	if err == nil {
		logger.SetOutput(file)
	}
	return logger
}
