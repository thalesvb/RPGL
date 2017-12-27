/*
Package logger provides a simplified interface for log.
*/
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
Logger contains available log types.
*/
type Logger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

/*
Map of loggers already instantiated.
*/
var mapLoggers = make(map[string]Logger)

/*
GetLogger returns a singleton instance for a prefixed logger.
*/
func GetLogger(prefix string) Logger {
	aLogger, found := mapLoggers[prefix]
	if !found {
		aLogger = initLogger(prefix, os.Stdout, os.Stdout, os.Stdout, os.Stdout)
		mapLoggers[prefix] = aLogger
	}
	return aLogger
}

func initLogger(
	prefix string,
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
) Logger {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	logger := Logger{
		Trace:   log.New(traceHandle, fmt.Sprintf("(%s) TRACE ", prefix), logFlags),
		Info:    log.New(infoHandle, fmt.Sprintf("(%s) INFO ", prefix), logFlags),
		Warning: log.New(warningHandle, fmt.Sprintf("(%s) WARNING ", prefix), logFlags),
		Error:   log.New(errorHandle, fmt.Sprintf("(%s) ERROR ", prefix), logFlags),
	}

	return logger
}
