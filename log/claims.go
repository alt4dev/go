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

// Group start a log group for the goroutine that calls this function.
// A group should be closed after. Use: `defer Claims{...}.Group(...).Close()`
func (claims Claims) Group(v ...interface{}) *GroupResult {
	title := fmt.Sprint(v...)
	return &GroupResult{
		logResult: service.Log(2, true, title, claims.parse(), proto.Log_NONE),
		claims: &claims,
	}
}

// Print send claims and the log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Print(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_NONE)
}

// Printf send claims and the log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Printf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_NONE)
}

// Println send claims and the log message to alt4. The log level is NONE. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Println(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_NONE)
}

// Info send claims and the log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Info(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_INFO)
}

// Infof send claims and the log message to alt4. The log level is INFO. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Infof(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_INFO)
}

// Infoln send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Infoln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_INFO)
}

// Debug send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Debug(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_DEBUG)
}

// Debugf send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Debugf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_DEBUG)
}

// Debugln send claims and the log message to alt4. The log level is DEBUG. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Debugln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_DEBUG)
}

// Warning send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Warning(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_WARNING)
}

// Warningf send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Warningf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_WARNING)
}

// Warningln send claims and the log message to alt4. The log level is WARNING. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Warningln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_WARNING)
}

// Error send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprint(a...)
func (claims Claims) Error(v ...interface{}) *service.LogResult {
	message := fmt.Sprint(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_ERROR)
}

// Errorf send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintf(a...)
func (claims Claims) Errorf(format string, v ...interface{}) *service.LogResult {
	message := fmt.Sprintf(format, v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_ERROR)
}

// Errorln send claims and the log message to alt4. The log level is ERROR. Log message will be formatted by fmt.Sprintln(a...)
func (claims Claims) Errorln(v ...interface{}) *service.LogResult {
	message := fmt.Sprintln(v...)
	return service.Log(2, false, message, claims.parse(), proto.Log_ERROR)
}

// Fatal This is equivalent to calling Print followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Fatal(v ...interface{}) {
	message := fmt.Sprint(v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInExit(1)
}

// Fatalf This is equivalent to calling Printf followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Fatalf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInExit(1)
}

// Fatalln This is equivalent to calling Println followed by os.Exit(1). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Fatalln(v ...interface{}) {
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInExit(1)
}

// Panic This is equivalent to calling Print followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Panic(v ...interface{}) {
	message := fmt.Sprint(v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInPanic(message)
}

// Panicf This is equivalent to calling Printf followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Panicf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInPanic(message)
}

// Panicln This is equivalent to calling Println followed by panic(). The log level is FATAL.
// This method will wait for the write to complete
func (claims Claims) Panicln(v ...interface{}) {
	message := fmt.Sprintln(v...)
	service.Log(2, false, message, claims.parse(), proto.Log_FATAL).Result()
	BuiltInPanic(message)
}

func (claims Claims) parse() []*proto.Claim {
	protoClaims := make([]*proto.Claim, 0)
	for key, i := range claims {
		var claimValue string
		var claimType proto.Claim_Type
		switch i.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			claimType = proto.Claim_NUMBER
			claimValue = fmt.Sprint(i)
		case float32, float64:
			claimType = proto.Claim_NUMBER
			claimValue = fmt.Sprint(i)
		case bool:
			claimType = proto.Claim_BOOLEAN
			claimValue = fmt.Sprint(i.(bool))
		case string:
			claimType = proto.Claim_STRING
			claimValue = i.(string)
		case time.Time:
			claimType = proto.Claim_TIMESTAMP
			claimValue = fmt.Sprint(i.(time.Time).UnixNano())
		default:
			claimType = proto.Claim_STRING
			claimValue = fmt.Sprint(i)
		}
		protoClaims = append(protoClaims, &proto.Claim{
			Name:     key,
			Type: claimType,
			Value:    claimValue,
		})
	}
	return protoClaims
}
