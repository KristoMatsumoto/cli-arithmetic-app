package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// Output to file
	file, err := os.OpenFile("app/log/output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Out = os.Stdout
		Log.Warn("Failed to log to file, using default stderr")
	} else {
		Log.Out = file
	}

	// Set default log level
	Log.SetLevel(logrus.DebugLevel)
}
