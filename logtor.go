// Package logtor provides a flexible logging framework that allows the coordination of multiple log creators
// with different destinations and log levels. It includes a central manager, Logtor, for managing log creators
// and controlling the global log level.
//
// Logtor allows you to log messages to various destinations simultaneously (e.g., file, console) and dynamically
// switch between different log creators. Each log creator must implement the LogCreator interface, providing
// methods for logging messages, retrieving the log creator's name, setting call depth, and performing cleanup
// operations during shutdown.
//
// Usage:
// - Create a new Logtor instance with NewLogtor().
// - Add log creators using AddLogCreators(), specifying destinations such as files or brokers.
// - Change the active log creator with ChangeLogCreator() to direct log messages to a specific log creator.
// - Set the global log level with SetLogLevel() to control which log messages are recorded.
// - Use LogIt() or LogItWithCallDepth() to log messages with the currently active log creator.
// - Gracefully shut down log creators using Shutdown().
package logtor

import (
	"reflect"
	"sync"

	"github.com/Eyup-Devop/logtor/types"
)

var defaultCreatorName string = "defaultCreator"

// New creates a new Logtor instance with default settings.
//
// It initializes a Logtor with an empty list of log creators, a global log level set to NONE,
// and no current log creator selected.
//
// Returns:
//   - *Logtor: A pointer to the newly created Logtor.
func New() *Logtor {
	return &Logtor{
		logCreatorList:    make(map[types.LogCreatorName]LogCreator),
		logLevel:          types.NONE,
		currentLogCreator: nil,
	}
}

func (l *Logtor) WithDefaultCreator(creator LogCreator) *Logtor {
	l.defaultCreator = creator
	return l
}

// Logtor is a central logging manager that coordinates multiple log creators and log levels.
//
// It manages a list of log creators, allowing you to log messages to different destinations (e.g., file, console) simultaneously.
// You can set the global log level for Logtor to control which log messages are recorded.
//
// Fields:
//   - logCreatorList: A map of LogCreatorName to LogCreator, representing registered log creator.
//   - logLevel: The global log level that controls which log messages are created.
//   - currentLogCreator: The currently active log creator for logging messages.
//   - changeMutex: A read-write mutex to control concurrent access to Logtor's fields.
type Logtor struct {
	logCreatorList    map[types.LogCreatorName]LogCreator
	logLevel          types.LogLevel
	currentLogCreator LogCreator
	changeMutex       sync.RWMutex
	defaultCreator    LogCreator
}

// SetLogLevel sets the global log level for the Logtor instance.
//
// You can use this method to change the log level for the Logtor, which controls which log messages
// are recorded and displayed. The log level should be one of the predefined LogLevelType constants.
//
// Parameters:
//   - logLevel: The new global log level to set for the Logtor.
func (l *Logtor) SetLogLevel(logLevel types.LogLevel) bool {
	if logLevel.IsValid() {
		l.logLevel = logLevel
		return true
	}
	return false
}

// LogLevel returns the current global log level of the Logtor instance.
//
// Use this method to retrieve the current global log level, which determines which log messages
// are recorded or displayed. The returned value is of type LogLevelType.
//
// Returns:
//   - LogLevelType: The current global log level.
func (l *Logtor) LogLevel() types.LogLevel {
	return l.logLevel
}

// ChangeLogCreator changes the active log creator to the one with the specified name.
//
// Use this method to switch the active log creator to the one identified by the provided
// LogCreatorName. This allows you to direct log messages to a specific log creator from the
// list of registered log creators.
//
// Parameters:
//   - logCreatorName: The name of the log creator to make active.
//
// Returns:
//   - bool: True if the log creator with the specified name exists and is successfully set as active;
//     false if the log creator does not exist.
func (l *Logtor) ChangeLogCreator(logCreatorName types.LogCreatorName) bool {
	l.changeMutex.RLock()
	defer l.changeMutex.RUnlock()
	if _, ok := l.logCreatorList[logCreatorName]; !ok {
		return false
	}
	l.currentLogCreator = l.logCreatorList[logCreatorName]
	return true
}

// LogCreator returns the currently active log creator of the Logtor instance.
//
// Use this method to obtain the currently active log creator, which is responsible for recording
// log messages at the global log level. The returned value is of type LogCreator.
//
// Returns:
//   - LogCreator: The currently active log creator.
func (l *Logtor) LogCreator() LogCreator {
	return l.currentLogCreator
}

// LogIt logs a message at the specified log level using the currently active log creator.
//
// This method allows you to log a message at a specific log level, subject to the global log level
// configured for the Logtor. If the provided log level is acceptable based on the global log level,
// the message is recorded by the currently active log creator.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: True if the message was successfully logged; false if it was skipped due to the log level.
func (l *Logtor) LogIt(level types.LogLevel, logMessage interface{}) bool {
	if l.logLevel.IsLogLevelAcceptable(level) && l.currentLogCreator.IsReady() {
		return l.currentLogCreator.LogIt(level, logMessage)
	} else if l.logLevel.IsLogLevelAcceptable(level) && !l.currentLogCreator.IsReady() && l.defaultCreator != nil {
		return l.defaultCreator.LogIt(level, logMessage)
	}
	return false
}

// LogIt logs a message at the specified log level using the currently active log creator.
//
// This method allows you to log a message at a specific log level, subject to the global log level
// configured for the Logtor. If the provided log level is acceptable based on the global log level,
// the message is recorded by the currently active log creator.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - callDepth: The call depth for calling function.
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: True if the message was successfully logged; false if it was skipped due to the log level.
func (l *Logtor) LogItWithCallDepth(level types.LogLevel, callDepth int, logMessage interface{}) bool {
	if types.IsLogLevelAcceptable(l.LogLevel(), level) && l.currentLogCreator.IsReady() {
		return l.currentLogCreator.LogItWithCallDepth(level, callDepth, logMessage)
	} else if l.logLevel.IsLogLevelAcceptable(level) && !l.currentLogCreator.IsReady() && l.defaultCreator != nil {
		return l.defaultCreator.LogItWithCallDepth(level, callDepth, logMessage)
	}
	return false
}

// AddLogcreators registers one or more log creators with the Logtor instance.
//
// This method allows you to add multiple log creators to the Logtor. The log creators are
// identified by their names and can be used for logging messages. If no active log creator
// is currently set, the first added log creator becomes the active one.
//
// Parameters:
//   - logCreators: One or more LogCreator instances to be added to the Logtor.
func (l *Logtor) AddLogCreators(logCreators ...LogCreator) {
	l.changeMutex.Lock()
	for _, logCreator := range logCreators {
		if logCreator != nil && !reflect.ValueOf(logCreator).IsNil() {
			l.logCreatorList[logCreator.LogName()] = logCreator
		}
	}
	l.changeMutex.Unlock()
	if l.currentLogCreator == nil {
		l.ChangeLogCreator(logCreators[0].LogName())
	}
}

// Shutdown gracefully shuts down all registered log creators.
//
// Use this method to perform any necessary cleanup or shutdown operations for all registered log creators.
// It iterates through the list of log creators and calls their respective shutdown methods.
func (l *Logtor) Shutdown() {
	for _, logCreator := range l.logCreatorList {
		logCreator.Shutdown()
	}
}
