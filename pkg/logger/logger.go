package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	prettyPrint := false
	if os.Getenv("DEBUG") == "true" {
		prettyPrint = true
	}

	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		PrettyPrint:       prettyPrint,
		DisableHTMLEscape: true,
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "datetime",
			log.FieldKeyLevel: "type",
			log.FieldKeyMsg:   "message",
			log.FieldKeyFunc:  "function",
			log.FieldKeyFile:  "file",
		},
	})
}

func makeMessage(inputs []interface{}) string {
	if len(inputs) == 0 {
		return "(MISSING)"
	}
	var message string
	for _, x := range inputs {
		if message != "" {
			message += " "
		}
		message += fmt.Sprint(x)
	}
	return message
}

func getPathFile(file string) string {
	dir, _ := os.Getwd()
	filePath, err := filepath.Rel(dir, file)
	if err != nil {
		panic(err)
	}
	return filePath
}

func LogInfo(inputs ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Info(makeMessage(inputs))
}

func LogWarn(inputs ...interface{}) {
	message := makeMessage(inputs)
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(message)
}

func LogErr(inputs ...interface{}) {
	message := makeMessage(inputs)
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	log.WithFields(log.Fields{
		"file": file,
		"line": line,
	}).Error(message)
}

// Parameters:
//   - userId: A string representing the user ID. This is a unique identifier for the user performing the action being loged.
//   - data: An interface{} representing the request data. This could be any data coming through the request such as request parameters, request body, etc.
//   - inputs: first element: the first parameter will be the response of a function or process in question
//   - A variadic number of interface{} representing the messages to log.
func LogInfoWithDetails(userId string, data interface{}, inputs ...interface{}) {
	var response interface{}
	if len(inputs) > 1 {
		response = inputs[0]
		inputs = inputs[1:]
	}

	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)

	fields := log.Fields{
		"userId":    userId,
		"file":      file,
		"line":      line,
		"inputData": data,
	}

	if response != nil {
		fields["response"] = response
	}

	log.WithFields(fields).Info(makeMessage(inputs))
}

// Parameters:
// - userId: A string representing the user ID. This is a unique identifier for the user performing the action being loged.
// - data: An interface{} representing the request data. This could be any data coming through the request such as request parameters, request body, etc.
// - inputs: A variadic number of interface{} representing the messages to log.
func LogErrWithDetails(userId string, data interface{}, inputs ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	log.WithFields(log.Fields{
		"userId": userId,
		"file":   file,
		"line":   line,
		"data":   data,
	}).Error(makeMessage(inputs))
}
