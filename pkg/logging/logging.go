package logging

import (
	"fmt"
	"os"

	restfullog "github.com/emicklei/go-restful/log"
	log "github.com/sirupsen/logrus"
)

func getLogLevel(logLevel string) log.Level {
	switch logLevel {
	case "Error":
		return log.ErrorLevel
	case "Warn":
		return log.WarnLevel
	case "Info":
		return log.InfoLevel
	case "Debug":
		return log.DebugLevel
	case "Trace":
		return log.TraceLevel
	}
	fmt.Printf("invalid log level")
	return log.InfoLevel
}

// InitializeLogger initialized the logger and restful logger used by the server.
func InitializeLogger(logDir string, logFile string, logLevel string) {
	formatter := new(log.TextFormatter)

	// Set timestamp format
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.FullTimestamp = true

	// Set the formatter
	log.SetFormatter(formatter)

	// Setup the restful logger
	restfulLogger := log.New()
	restfulLogger.Formatter = formatter
	restfullog.SetLogger(restfulLogger)

	// Set the log file
	f, err := os.OpenFile(logDir+logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
		restfulLogger.Out = f
	}

	// Set the log level
	level := getLogLevel(logLevel)
	log.SetLevel(level)
	restfulLogger.Level = level
}
