package loggers

import (
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Entry

const (
	AppName = "s3-file-handler"
)

type LogFields struct {
	FileName   string `json:"fileName"`
	Identifier string `json:"identifier"`
	Action     string `json:"action"`
	StatusCode int    `json:"statusCode,omitempty"`
}

//var debug config.Config.Debug

func init() {
	//set properties for the logging

	//	log.SetFormatter(&log.JSONFormatter{}, JSONPretty:true) //prints pretty
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Set application name that will be include in all log messages
	Logger = log.WithFields(log.Fields{
		"application": AppName,
	})
}

// DebugWithInputData Function to log a DEBUG level log message that includes the input data
func DebugWithInputData(inputName string, input ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file":  path.Base(fileName),
		"line":  lineNumber,
		"input": inputName,
	}).Debug(input...)
}

// DebugWithResultData Function to log a DEBUG level log message that includes the result data
func DebugWithResultData(resultName string, input ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file":   path.Base(fileName),
		"line":   lineNumber,
		"result": resultName,
	}).Debug(input...)
}

// Debug Function to log a DEBUG level log message
func Debug(logFields LogFields, logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file":       path.Base(fileName),
		"line":       lineNumber,
		"fileName":   logFields.FileName,
		"identifier": logFields.Identifier,
	}).Info(logMsg...)
}

func Info(logFields LogFields, logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file":       path.Base(fileName),
		"line":       lineNumber,
		"fileName":   logFields.FileName,
		"identifier": logFields.Identifier,
	}).Info(logMsg...)
}

func Error(logFields LogFields, logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file":       path.Base(fileName),
		"line":       lineNumber,
		"fileName":   logFields.FileName,
		"identifier": logFields.Identifier,
	}).Error(logMsg...)
}

// Warn Function to log a WARNING level log message
func Warn(logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file": path.Base(fileName),
		"line": lineNumber,
	}).Warn(logMsg...)
}

// Fatal Function to log a Fatal level log message
func Fatal(logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file": path.Base(fileName),
		"line": lineNumber,
	}).Fatal(logMsg...)
}

// Panic Function to log a PANIC level log message
func Panic(logMsg ...interface{}) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	Logger.WithFields(log.Fields{
		"file": path.Base(fileName),
		"line": lineNumber,
	}).Panic(logMsg...)
}
