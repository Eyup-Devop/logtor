package logtor

import "github.com/Eyup-Devop/logtor/types"

// LogCreator is an interface for log creator that handle log messages of different log levels.
//
// A log creator must implement the following methods:
//
// - LogIt: Logs a message with the specified log level.
// - LogName: Returns the name of the log creator.
// - Shutdown: Performs any necessary cleanup or shutdown operations for the log creator.
type LogCreator interface {
	// LogIt logs a message with the specified log level and returns true if successful.
	LogIt(level types.LogLevel, logMessage interface{}) bool

	// LogItWithCallDepth logs a message with the specified log level and call depth and returns true if successful.
	LogItWithCallDepth(level types.LogLevel, callDepth int, logMessage interface{}) bool

	// LogName returns the name of the log creators.
	LogName() types.LogCreatorName

	// SetCallDepth sets the call depth for the log creator.
	SetCallDepth(callDepth int)

	// CallDepth returns the call depth for the log creator.
	CallDepth() int

	// IsReady() returns true if the log creator is ready to log messages.
	IsReady() bool

	// Shutdown performs any necessary cleanup or shutdown operations for the log creator.
	Shutdown()
}
