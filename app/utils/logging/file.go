package logging

import (
	"fmt"
	"gin-lib/conf/setting"
	"log"
	"os"
	"time"
)

var (
	// LogSavePath = "runtime/logs/"
	LogSavePath = setting.Conf().Log.Path
	// LogSaveName = "log"
	LogSaveName = setting.Conf().Log.Prefix
	// SQLLogSaveName SqlLogSaveName
	SQLLogSaveName = setting.Conf().Log.SQLPrefix
	// LogFileExt log file ext
	LogFileExt = "log"
	// TimeFormat time format
	TimeFormat = "2006-01-02"
)

// getLogFilePath return the log file path
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// getLogFilePath return the log file full path
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// getSqlLogFile return the sql log file
func getSQLLogFile() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", SQLLogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile open file and return file source object
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

// SQLLog SQLLog io.writer
func SQLLog() *log.Logger {
	sqlPath := getSQLLogFile()
	S := openLogFile(sqlPath)
	return log.New(S, "\r\n", 7|8)
}

// mkDir create log file
func mkDir() {
	// dir, _ := os.Getwd()
	// err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	err := os.MkdirAll(getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
