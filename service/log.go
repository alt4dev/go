package service

import (
	"fmt"
	"github.com/alt4dev/protobuff/proto"
	"runtime"
	"strings"
	"time"
)

// Log Creates a log entry and writes it to alt4 in the background.
// This function should not be called directly and should instead be used from helper functions under the `log` package.
func Log(calldepth int, asGroup bool, message string, claims []*proto.Claim, level proto.Log_Level, logTime time.Time) *LogResult {
	if asGroup {
		initGroup()
	}
	// Get the parent file and function of the caller
	pc, file, line, _ := runtime.Caller(calldepth)
	function := runtime.FuncForPC(pc).Name()
	msg := proto.Log{
		Source:    options.Source,
		Thread:    getThreadId(),
		Message:   message,
		Claims:    claims,
		File:      file,
		Line:      uint32(line),
		Function:  function,
		Level:     level,
		Timestamp: uint64(logTime.UnixNano()),
		Group:     asGroup,
	}
	result := LogResult{
		wg: WaitGroup(),
	}
	if options.Mode == ModeDebug || options.Mode == ModeTesting {
		// Write to stderr if conditions are met.
		emitLog(&msg)
	}
	if options.Mode != ModeTesting && options.Mode != ModeSilent {
		WaitGroup().Add(1)
		go writerHelper(&msg, &result)
	}
	return &result
}

func writerHelper(msg *proto.Log, result *LogResult) {
	defer result.wg.Done()
	Alt4RemoteHelper.WriteLog(msg, result)
}

func emitLog(msg *proto.Log) {
	timeString := time.Unix(0, int64(msg.Timestamp)).Format("2006-01-02T15:04:05.000Z")
	message := fmt.Sprintf("[alt4 %s] %s %s:%d %s", msg.Level.String(), timeString, msg.File, msg.Line, msg.Message)
	lines := []string{message}
	for _, claim := range msg.Claims {
		lines = append(lines, fmt.Sprintf("\tclaim.%s: '%s'", claim.Name, claim.Value))
	}
	emit.Println(strings.Join(lines, "\n"))
}

