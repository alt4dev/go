package log

import (
	"fmt"
	"github.com/alt4dev/go/service"
	"github.com/alt4dev/protobuff/proto"
	"time"
)

// Claims are fields that will can be associated to your log entry.
// They can be used to filter and better identify your logs.
type Claims map[string]interface{}

func (claims Claims) parse() []*proto.Claim {
	return ParseClaims(claims)
}

// Group start a log group for the goroutine that calls this function.
// A group should be closed after. Use: `defer Claims{...}.Group(...).Close()`
func (claims Claims) Group(v ...interface{}) *GroupResult {
	title := fmt.Sprint(v...)
	return &GroupResult{
		logResult: service.Log(2, true, title, nil, LEVEL.DEBUG),
	}
}

// Print send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Print(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), LEVEL.DEBUG)
}

// Printf send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Printf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), LEVEL.DEBUG)
}

// Println send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Println(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), LEVEL.DEBUG)
}

// Info send claims and the log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Info(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, LEVEL.INFO)
}

// Infof send claims and the log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Infof(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, LEVEL.INFO)
}

// Infoln send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Infoln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, LEVEL.INFO)
}

// Debug send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Debug(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, LEVEL.DEBUG)
}

// Debugf send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Debugf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, LEVEL.DEBUG)
}

// Debugln send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Debugln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, LEVEL.DEBUG)
}

// Warning send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Warning(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, LEVEL.WARNING)
}

// Warningf send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Warningf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, LEVEL.WARNING)
}

// Warningln send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Warningln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, LEVEL.WARNING)
}

// Error send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Error(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, nil, LEVEL.ERROR)
}

// Errorf send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Errorf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, nil, LEVEL.ERROR)
}

// Errorln send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Errorln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, nil, LEVEL.ERROR)
}

// Fatal This is equivalent to calling Print followed by os.Exit(1). The log level is CRITICAL.
func (claims Claims) Fatal(v ...interface{}) {
	message := fmt.Sprint(v...)
	service.Log(2, false, message, nil, LEVEL.ERROR).Result()
	BuiltInExit(1)
}

// Fatalf This is equivalent to calling Printf followed by os.Exit(1). The log level is CRITICAL.
func (claims Claims) Fatalf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, nil, LEVEL.ERROR).Result()
	BuiltInExit(1)
}

// Fatalln This is equivalent to calling Println followed by os.Exit(1). The log level is CRITICAL.
func (claims Claims) Fatalln(v ...interface{}) {
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, nil, LEVEL.ERROR).Result()
	BuiltInExit(1)
}

// Panic This is equivalent to calling Print followed by panic(). The log level is CRITICAL.
func (claims Claims) Panic(v ...interface{}) {
	message := fmt.Sprint(v...)
	service.Log(2, false, message, nil, LEVEL.CRITICAL).Result()
	BuiltInPanic(message)
}

// Panicf This is equivalent to calling Printf followed by panic(). The log level is CRITICAL.
func (claims Claims) Panicf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, nil, LEVEL.CRITICAL).Result()
	BuiltInPanic(message)
}

// Panicln This is equivalent to calling Println followed by panic(). The log level is CRITICAL.
func (claims Claims) Panicln(v ...interface{}) {
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, nil, LEVEL.CRITICAL).Result()
	BuiltInPanic(message)
}

// ParseClaims This function is used internally to convert Claims to a format that can be sent to alt4
func ParseClaims(claims Claims) []*proto.Claim {
	protoClaims := make([]*proto.Claim, 0)
	for key, i := range claims {
		var claimValue string
		var claimType uint8
		switch i.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			claimType = 1
			claimValue = fmt.Sprint(i)
		case float32, float64:
			claimType = 2
			claimValue = fmt.Sprint(i)
		case bool:
			claimType = 3
			claimValue = fmt.Sprint(i.(bool))
		case string:
			claimType = 4
			claimValue = i.(string)
		case time.Time:
			claimType = 5
			claimValue = fmt.Sprint(i.(time.Time).UnixNano())
		default:
			claimType = 0
			claimValue = fmt.Sprint(i)
		}
		protoClaims = append(protoClaims, &proto.Claim{
			Name:     key,
			DataType: uint32(claimType),
			Value:    claimValue,
		})
	}
	return protoClaims
}
