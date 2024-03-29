package applogger

var logger Logger = &DefaultLogger{level: DebugLevel}

// SetLogger sets the current logger.
func SetLogger(l Logger) {
	logger = l
}

// GetLogger returns the current logger.
func GetLogger() Logger {
	return logger
}

func Error(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}
