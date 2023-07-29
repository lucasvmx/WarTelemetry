package logger

import (
	"fmt"
	"log"
	"os"
)

var debug bool = false
var prefix string = "wartelemetry"
var internalLogger *log.Logger

func EnableDebug() {
	internalLogger = log.New(os.Stdout, "", log.Lshortfile)
	internalLogger.SetFlags(0)
	debug = true
}

func DisableLogger() {
	debug = false
}

func LogInfo(format string, args ...any) {
	if debug {
		fmt.Printf("%v:INFO ", prefix)
		internalLogger.Printf(format, args...)
	}
}

func LogError(format string, args ...any) {
	if debug {
		fmt.Printf("%v:ERROR ", prefix)
		internalLogger.Printf(format, args...)
	}
}

func LogSuccess(format string, args ...any) {
	if debug {
		fmt.Printf("%v:SUCCESS ", prefix)
		internalLogger.Printf(format, args...)
	}
}

func LogCritical(format string, args ...any) {
	if debug {
		fmt.Printf("%v:CRITICAL ", prefix)
		internalLogger.Printf(format, args...)
	}
}
