package logs

import "github.com/sirupsen/logrus"

var Logs = logrus.New()

func InitLoggers() {
	jsonFormatter := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	}

	Logs.SetFormatter(&jsonFormatter)
	Logs.SetReportCaller(true)
	Logs.SetLevel(logrus.InfoLevel)
}
