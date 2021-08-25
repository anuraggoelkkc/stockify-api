package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"stockify-api/src/constants"
)

func InitializeLogging(logFile, env string) {
	dir := path.Dir(logFile)
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}
	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Println("Could Not Open Log File : " + err.Error())
		return
	}

	logrus.SetOutput(file)
	if env == constants.EnvProduction {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	//logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
