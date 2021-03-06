// Package log provides methods for writing logs to Alt4
// Methods in this package write logs asynchronously unless otherwise specified.
// You can call the function `Result` after log which will block if the operation is not done.
// The advised behaviour is to group your logs and defer the Close method which will wait for all operations to finish
package log

import (
	"fmt"
	"github.com/alt4dev/go/service"
	"github.com/alt4dev/protobuff/proto"
	"os"
)

// BuiltInPanic Internally this function just calls panic(). Override for testing(Panic, Panicf, Panicln)
var BuiltInPanic func(v interface{}) = func(v interface{}) {
	panic(v)
}

// BuiltInExit Internally this function just calls os.Exit. Override for testing(Fatal, Fatalf, Fatalln)
var BuiltInExit func(code int) = func(code int) {
	os.Exit(code)
}

// Group start a log group for the goroutine that calls this function.
// A group should be closed after. Use: `defer Group(...).Close()`
func Group(v ...interface{}) *GroupResult {
	t := service.LogTime()
	title := fmt.Sprint(v...)
	return &GroupResult{
		logResult: service.Log(2, true, title, nil, proto.Log_NONE, t),
		claims: nil,
	}
}

// Print send a log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprint(a...)
func Print(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, proto.Log_NONE, t)
}

// Printf send a log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprintf(a...)
func Printf(format string, v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, proto.Log_NONE, t)
}

// Println send a log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprintln(a...)
func Println(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, proto.Log_NONE, t)
}

// Info send a log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprint(a...)
func Info(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, proto.Log_INFO, t)
}

// Infof send a log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprintf(a...)
func Infof(format string, v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, proto.Log_INFO, t)
}

// Infoln send a log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func Infoln(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, proto.Log_INFO, t)
}

// Debug send a log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprint(a...)
func Debug(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, proto.Log_DEBUG, t)
}

// Debugf send a log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintf(a...)
func Debugf(format string, v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, proto.Log_DEBUG, t)
}

// Debugln send a log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func Debugln(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, proto.Log_DEBUG, t)
}

// Warning send a log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprint(a...)
func Warning(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, proto.Log_WARNING, t)
}

// Warningf send a log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintf(a...)
func Warningf(format string, v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, proto.Log_WARNING, t)
}

// Warningln send a log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintln(a...)
func Warningln(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, proto.Log_WARNING, t)
}

// Error send a log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprint(a...)
func Error(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, proto.Log_ERROR, t)
}

// Errorf send a log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintf(a...)
func Errorf(format string, v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, proto.Log_ERROR, t)
}

// Errorln send a log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintln(a...)
func Errorln(v ...interface{}) *service.LogResult {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, proto.Log_ERROR, t)
}

// Fatal This is equivalent to calling Print followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func Fatal(v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInExit(1)
}

// Fatalf This is equivalent to calling Printf followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func Fatalf(format string, v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInExit(1)
}

// Fatalln This is equivalent to calling Println followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func Fatalln(v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInExit(1)
}

// Panic This is equivalent to calling Print followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func Panic(v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprint(v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInPanic(message)
}

// Panicf This is equivalent to calling Printf followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func Panicf(format string, v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInPanic(message)
}

// Panicln This is equivalent to calling Println followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func Panicln(v ...interface{}) {
	t := service.LogTime()
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, nil, proto.Log_FATAL, t).Result()
	BuiltInPanic(message)
}
