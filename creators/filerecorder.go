// Package logtor provides log creators and loggers for various destinations.
//
// It includes implementations for creating logs and logging messages to a file,
// with options for customizing log formatting and output location.
package creators

import (
	"fmt"
	"log"
	"os"

	"github.com/Eyup-Devop/logtor"
	"github.com/Eyup-Devop/logtor/types"
)

// NewFileCreator creates a new instance of FileCreator, which logs messages to a file.
//
// It initializes a FileCreator with the provided file name, log creator name, call depth, and log prefix.
//
// Parameters:
//   - filename: The name of the log file.
//   - logName: The name representing the log creator (e.g., File).
//   - callDepth: The call depth to be used in log output.
//   - logPrefix: An integer representing log prefix settings.
//
// Returns:
//   - *FileCreator: A pointer to the newly created FileCreator.
//   - error: An error if initialization fails, or nil if successful.
//
// If logName is an empty string, it defaults to File.
func NewFileCreator(filename string, logName types.LogCreatorName, callDepth int, logPrefix int) (logtor.LogCreator, error) {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}

	fileCreator := &FileCreator{
		log:       log.New(logFile, "", log.LstdFlags|log.Lshortfile),
		fileName:  filename,
		logName:   logName,
		callDepth: callDepth,
		logPrefix: logPrefix,
	}
	// Set default log name if not provided
	if logName == "" {
		fileCreator.logName = File
	}

	return fileCreator, nil
}

// File is a constant representing the LogCreatorName for the File log creator.
const File types.LogCreatorName = "File"

// FileCreator is an implementation of the LogCreator interface for logging messages to a file.
type FileCreator struct {
	log       *log.Logger
	fileName  string
	logName   types.LogCreatorName
	callDepth int
	logPrefix int
}

// LogItWithCallDepth logs a message with the specified log level, call depth, and log message to the file.
//
// It formats the log entry with the log level's prefix and then outputs the log message.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - callDepth: The call depth for recording the log entry.
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (fr *FileCreator) LogItWithCallDepth(level types.LogLevel, callDepth int, logMessage interface{}) bool {
	fr.log.SetPrefix(fmt.Sprintf("%-*s : ", fr.logPrefix, level))
	fr.log.Output(callDepth, fmt.Sprintf("%+v", logMessage))
	return true
}

// LogIt logs a message with the specified log level using the default call depth to the file.
//
// This method is a convenience wrapper around LogItWithCallDepth, using the call depth
// configured for the FileCreator instance.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (fr *FileCreator) LogIt(level types.LogLevel, logMessage interface{}) bool {
	return fr.LogItWithCallDepth(level, fr.callDepth, logMessage)
}

// LogName returns the name of the log creator.
//
// Returns:
//   - LogCreatorName: The name of the log creator.
func (fr *FileCreator) LogName() types.LogCreatorName {
	return fr.logName
}

// SetCallDepth sets the call depth for recording log entries.
//
// This method allows configuring how deep into the call stack the logger should trace when recording
// log messages. A higher call depth includes more layers of function calls in the log output,
// providing additional context about the log origin.
//
// Parameters:
//   - callDepth: The depth to set for recording log entries.
func (fr *FileCreator) SetCallDepth(callDepth int) {
	fr.callDepth = callDepth
}

// CallDepth returns the current call depth setting for recording log entries.
//
// Returns:
//   - int: The current call depth setting for recording log entries.
func (fr *FileCreator) CallDepth() int {
	return fr.callDepth
}

// Shutdown performs any necessary cleanup or shutdown operations for the log creator.
//
// This method is present to satisfy the LogCreator interface, but it does not perform any actions
// in the case of the FileCreator. It is left empty intentionally.
func (fr *FileCreator) Shutdown() {
	// No cleanup or shutdown actions needed for FileCreator.
}

func (fr *FileCreator) IsReady() bool {
	return true
}
