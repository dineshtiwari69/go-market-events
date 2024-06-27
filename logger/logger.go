package logger

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func SetLogger(logLevel string) {
	//set log level
	customFormatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(fmt.Errorf("invalid log level: %s", logLevel))
	}

	logrus.SetLevel(lvl)
}
