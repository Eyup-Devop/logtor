package creators

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Eyup-Devop/logtor"
	"github.com/Eyup-Devop/logtor/types"
)

// NewBaseCreator creates a new instance of the BaseCreator.
//
// It initializes a BaseCreator with the specified logName, callDepth, and logPrefix.
//
// Parameters:
//   - logName: The type of log creator (e.g., File, Console).
//   - callDepth: The call depth to be used in log output.
//   - logPrefix: An integer representing log prefix settings.
//
// Returns:
//   - *BaseCreator: A pointer to the newly created BaseCreator.
//   - error: An error if initialization fails, or nil if successful.
//
// If logName is an empty string, it defaults to Console.
func NewBaseCreator(logName types.LogCreatorName, callDepth int, logPrefix int) (logtor.LogCreator, error) {
	baseCreator := &BaseCreator{
		log:       log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile),
		logName:   logName,
		callDepth: callDepth,
		logPrefix: logPrefix,
	}

	if logName == "" {
		baseCreator.logName = Console
	}

	return baseCreator, nil
}

// Console is a constant representing the LogCreatorName for the Console log creator.
const Console types.LogCreatorName = "Console"

// BaseCreator is a basic implementation of the LogCreator interface.
// It logs messages with a specified log level, call depth, and log prefix.
type BaseCreator struct {
	log       *log.Logger
	logName   types.LogCreatorName
	callDepth int
	logPrefix int
}

// LogItWithCallDepth logs a message with the specified log level, call depth, and log message.
//
// It formats the log entry with the log level's color, log prefix, and then outputs the log message.
// The call depth parameter determines how many stack frames to ascend when recording the log entry.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - callDepth: The call depth for recording the log entry.
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (br *BaseCreator) LogItWithCallDepth(level types.LogLevel, callDepth int, logMessage interface{}) bool {
	message, _ := json.Marshal(logMessage)
	br.log.SetPrefix(fmt.Sprintf("%s%-*s : ", types.GetColorForLogLevel(level), br.logPrefix, level))
	br.log.Output(callDepth, fmt.Sprintf("%+v%s", string(message), types.ResetColor))
	return true
}

// LogIt logs a message with the specified log level using the default call depth.
//
// This method is a convenience wrapper around LogItWithCallDepth, using the call depth
// configured for the BaseCreator instance.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (br *BaseCreator) LogIt(level types.LogLevel, logMessage interface{}) bool {
	return br.LogItWithCallDepth(level, br.callDepth, logMessage)
}

// LogName returns the name of the log creator.
//
// Returns:
//   - LogCreatorName: The name of the log creator.
func (br *BaseCreator) LogName() types.LogCreatorName {
	return br.logName
}

// SetCallDepth sets the call depth for recording log entries.
//
// This method allows configuring how deep into the call stack the logger should trace when recording
// log messages. A higher call depth includes more layers of function calls in the log output,
// providing additional context about the log origin.
//
// Parameters:
//   - callDepth: The depth to set for recording log entries.
func (br *BaseCreator) SetCallDepth(callDepth int) {
	br.callDepth = callDepth
}

// CallDepth returns the current call depth setting for recording log entries.
//
// Returns:
//   - int: The current call depth setting for recording log entries.
func (br *BaseCreator) CallDepth() int {
	return br.callDepth
}

// Shutdown performs any necessary cleanup or shutdown operations for the log creator.
//
// This method is present to satisfy the LogCreator interface, but it does not perform any actions
// in the case of the BaseCreator. It is left empty intentionally.
func (br *BaseCreator) Shutdown() {
	// No cleanup or shutdown actions needed for BaseCreator.
}
