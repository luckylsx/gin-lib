package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level error level
type Level int

var (
	// F file object source
	F *os.File

	// DefaultPrefix DefaultPrefix
	DefaultPrefix = ""
	// DefaultCallerDepth DefaultCallerDepth
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	// DEBUG DEBUG
	DEBUG Level = iota
	// INFO INFO
	INFO
	// WARNING WARNING
	WARNING
	// ERROR ERROR
	ERROR
	// FATAL FATAL
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug DEBUG level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info INFO level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn WARN level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error ERROR level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal FATAL level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// setPrefix set the log info prefix
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
