package logtor_test

import (
	"testing"
	"time"

	"github.com/Eyup-Devop/logtor"
	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

// TestLogtorUsingBaseCreatorWithString tests the functionality of Logtor with a base creator
// configured for string logging. It covers various log levels and call depths to ensure correct
// log message generation and handling.
//
// The test initializes a Logtor with a base creator configured for string logging and specific
// call depth settings. It sets the global log level to TRACE and logs messages at different levels,
// checking if each log entry is correctly generated. The test also logs messages with varying call
// depths to verify that the call depth parameter is considered during log creation.
//
// Test Steps:
//  1. Create a base log creator for string logging with a specified name, log level, and call depth.
//  2. Initialize a new Logtor and add the base creator to it.
//  3. Set the global log level of the Logtor to TRACE.
//  4. Log messages at different log levels using LogIt and LogItWithCallDepth methods.
//  5. Check if the log entries are generated as expected.
func TestLogtorUsingBaseCreatorWithString(t *testing.T) {
	baseCreator, err := creators.NewBaseCreator("Console", 3, 5)
	if err != nil {
		t.Error(err)
	}

	newLogtor := logtor.New()
	newLogtor.AddLogCreators(baseCreator)
	newLogtor.SetLogLevel(types.TRACE)

	newLogtor.LogIt(types.FATAL, "Example Test Log String")
	newLogtor.LogIt(types.ERROR, "Example Test Log String")
	newLogtor.LogIt(types.WARN, "Example Test Log String")
	newLogtor.LogIt(types.DEBUG, "Example Test Log String")
	newLogtor.LogIt(types.INFO, "Example Test Log String")
	newLogtor.LogIt(types.TRACE, "Example Test Log String")

	newLogtor.LogItWithCallDepth(types.FATAL, 0, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.ERROR, 1, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.WARN, 2, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.DEBUG, 3, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.INFO, 4, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.TRACE, 5, "Example Test Log String With Call Depth")
}

func TestLogtorUsingFileCreatorWithString(t *testing.T) {
	fileCreator, err := creators.NewFileCreator("./temp/temp.log", "File", 3, 5)
	if err != nil {
		t.Error(err)
	}

	newLogtor := logtor.New()
	newLogtor.AddLogCreators(fileCreator)
	newLogtor.SetLogLevel(types.TRACE)

	newLogtor.LogIt(types.FATAL, "Example Test Log String")
	newLogtor.LogIt(types.ERROR, "Example Test Log String")
	newLogtor.LogIt(types.WARN, "Example Test Log String")
	newLogtor.LogIt(types.DEBUG, "Example Test Log String")
	newLogtor.LogIt(types.INFO, "Example Test Log String")
	newLogtor.LogIt(types.TRACE, "Example Test Log String")

	newLogtor.LogItWithCallDepth(types.FATAL, 0, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.ERROR, 1, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.WARN, 2, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.DEBUG, 3, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.INFO, 4, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.TRACE, 5, "Example Test Log String With Call Depth")
}

func TestLogtorUsingBrokerCreatorWithString(t *testing.T) {
	brokers := []string{"127.0.0.1:19092"}
	brokerCreator, err := creators.NewBrokerCreator(brokers, "test", "Broker", 2, nil)
	if err != nil {
		t.Error(err)
	}

	newLogtor := logtor.New()
	newLogtor.AddLogCreators(brokerCreator)
	newLogtor.SetLogLevel(types.TRACE)

	newLogtor.LogIt(types.FATAL, "Example Test Log String")
	newLogtor.LogIt(types.ERROR, "Example Test Log String")
	newLogtor.LogIt(types.WARN, "Example Test Log String")
	newLogtor.LogIt(types.DEBUG, "Example Test Log String")
	newLogtor.LogIt(types.INFO, "Example Test Log String")
	newLogtor.LogIt(types.TRACE, "Example Test Log String")

	newLogtor.LogItWithCallDepth(types.FATAL, 0, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.ERROR, 1, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.WARN, 2, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.DEBUG, 3, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.INFO, 4, "Example Test Log String With Call Depth")
	newLogtor.LogItWithCallDepth(types.TRACE, 5, "Example Test Log String With Call Depth")

	time.Sleep(time.Second * 2)
}

func TestLogtorUsingAllCreators(t *testing.T) {
	baseCreator, err := creators.NewBaseCreator("Console", 3, 5)
	if err != nil {
		t.Error(err)
	}
	fileCreator, err := creators.NewFileCreator("./temp/temp.log", "File", 3, 5)
	if err != nil {
		t.Error(err)
	}
	brokers := []string{"127.0.0.1:19092"}
	brokerCreator, err := creators.NewBrokerCreator(brokers, "test", "Broker", 2, nil)
	if err != nil {
		t.Error(err)
	}

	newLogtor := logtor.New()
	newLogtor.AddLogCreators(baseCreator, fileCreator, brokerCreator)
	newLogtor.SetLogLevel(types.TRACE)

	newLogtor.ChangeLogCreator(creators.Console)
	if !newLogtor.LogIt(types.FATAL, "Example Test Log Console String") {
		t.Error("Failed to log to console")
	}

	newLogtor.ChangeLogCreator(creators.Broker)
	newLogtor.SetLogLevel(types.FATAL)
	if newLogtor.LogIt(types.ERROR, "Example Test Log Broker String") {
		t.Error("It suppose not to log it")
	}
	newLogtor.LogIt(types.WARN, "Example Test Log String")
	newLogtor.LogIt(types.DEBUG, "Example Test Log String")
	newLogtor.LogIt(types.INFO, "Example Test Log String")
	newLogtor.LogIt(types.TRACE, "Example Test Log String")
}
