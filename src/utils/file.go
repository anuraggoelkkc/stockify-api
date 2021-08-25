package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func GetFile(filePath string) (f *os.File, err error) {
	dir := path.Dir(filePath)
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}
	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorln("Unable to open file:", filePath)
	}
	return f, err
}
