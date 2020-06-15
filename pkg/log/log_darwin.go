package log

import (
	"log/syslog"

	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func GetLogLevel(logLevel string) log.Level {
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
	log.WithFields(log.Fields{
		"event": "getLogLevel",
		"topic": "logging",
		"key":   "logLevel",
	}).Warning("invalid log level")
	return log.WarnLevel
}

func InitializeLogger(logLevel log.Level) {
	formatter := new(log.TextFormatter)

	// Set timestamp format
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.FullTimestamp = true

	// Set the formatter
	log.SetFormatter(formatter)

	// Set the log level
	log.SetLevel(logLevel)

	// Uncomment if you want to log  the method name.   Note that this adds considerable overhead.
	// log.SetReportCaller(true)

	// Log to Syslogd
	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}
