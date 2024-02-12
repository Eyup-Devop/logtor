// Package logtor provides constants and utility functions related to logging, including LogLevel constants,
// LogCreatorName type, color codes for different log levels, and functions for color formatting and checking
// whether a log level is acceptable based on a selected level.
//
// Variables:
//   - ResetColor: ANSI escape code to reset color.
//   - NoneColor, FatalColor, ErrorColor, WarnColor, DebugColor, InfoColor, TraceColor:
//     ANSI escape codes for log level colors.
//
// Constants:
// - LogLevel: Represents different log levels (NONE, FATAL, ERROR, WARN, DEBUG, INFO, TRACE).
// - LogCreatorName: Represents the names of log creators.
// - Color Codes: ANSI escape codes for log level colors.
//
// Functions:
// - GetColorForLogLevel: Returns the ANSI escape code for the color associated with a log level.
// - IsLogLevelAcceptable: Checks if a given log level is acceptable based on the selected log level.
package types

type LogLevel string

const (
	NONE  LogLevel = "NONE"
	FATAL LogLevel = "FATAL"
	ERROR LogLevel = "ERROR"
	WARN  LogLevel = "WARN"
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	TRACE LogLevel = "TRACE"
)

var LogLevelList = []LogLevel{NONE, FATAL, ERROR, WARN, DEBUG, INFO, TRACE}

type LogCreatorName string

var (
	ResetColor = "\033[0m"
	NoneColor  = "\033[97m"
	FatalColor = "\033[31m"
	ErrorColor = "\033[31m"
	WarnColor  = "\033[33m"
	DebugColor = "\033[32m"
	InfoColor  = "\033[34m"
	TraceColor = "\033[35m"
)

func GetColorForLogLevel(level LogLevel) string {
	switch level {
	case FATAL:
		return FatalColor
	case ERROR:
		return ErrorColor
	case WARN:
		return WarnColor
	case DEBUG:
		return DebugColor
	case INFO:
		return InfoColor
	case TRACE:
		return TraceColor
	default:
		return ResetColor
	}
}

func IsLogLevelAcceptable(selected, using LogLevel) bool {
	switch selected {
	case FATAL:
		if using == FATAL {
			return true
		}
		return false
	case ERROR:
		if using == FATAL || using == ERROR {
			return true
		}
		return false
	case WARN:
		if using == FATAL || using == ERROR || using == WARN {
			return true
		}
		return false
	case DEBUG:
		if using == FATAL || using == ERROR || using == WARN || using == DEBUG {
			return true
		}
		return false
	case INFO:
		if using == FATAL || using == ERROR || using == WARN || using == DEBUG || using == INFO {
			return true
		}
		return false
	case TRACE:
		if using == FATAL || using == ERROR || using == WARN || using == DEBUG || using == INFO || using == TRACE {
			return true
		}
		return false
	default:
		return false
	}
}
