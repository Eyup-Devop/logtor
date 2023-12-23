// Package logtor_test provides test cases for the logtor package, specifically focusing on the BrokerCreator implementation.
//
// It includes test functions for logging various types of messages using the BrokerCreator,
// including string messages, structured data in a struct, and JSON-encoded data.
//
// Test Functions:
//   - TestBrokerCreatorWithString: Tests logging a string message with the BrokerCreator at the ERROR level.
//   - TestBrokerCreatorWithStruct: Tests logging a struct with the BrokerCreator at the WARN and INFO levels.
//   - TestBrokerCreatorWithJson: Tests logging JSON-encoded data with the BrokerCreator at the DEBUG and TRACE levels.
//
// These tests require a Kafka broker running locally on 127.0.0.1:19092 and may take a few seconds to complete due to sleep periods.
// Adjust the broker address and sleep durations based on your specific Kafka setup and test requirements.
package creators_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

var brokers = []string{"127.0.0.1:19092"}

// TestBrokerCreatorWithString tests logging a string message with the BrokerCreator at the ERROR level.
//
// It initializes a BrokerCreator with specified settings, logs an example string message at the ERROR level,
// and checks if the log entry is recorded successfully. The test includes a sleep period to allow time for message processing.
func TestBrokerCreatorWithString(t *testing.T) {
	brokerCreator, err := creators.NewBrokerCreator(brokers, "test", "Broker", 2, nil)
	if err != nil {
		t.Error(err)
	}
	if result := brokerCreator.LogIt(types.ERROR, "Example Log Message"); !result {
		t.Error("Log not recorded")
	}
	time.Sleep(time.Second * 2)
	brokerCreator.Shutdown()
}

// TestBrokerCreatorWithStruct tests logging a struct with the BrokerCreator at the WARN and INFO levels.
//
// It initializes a BrokerCreator with specified settings, creates an example struct,
// logs the struct at the WARN and INFO levels, and checks if the log entries are recorded successfully.
// The test includes a sleep period to allow time for message processing.
func TestBrokerCreatorWithStruct(t *testing.T) {
	brokerCreator, err := creators.NewBrokerCreator(brokers, "test", "Broker", 2, nil)
	if err != nil {
		t.Error(err)
	}

	exampleStruct := &struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Example Name",
		Age:  25,
	}

	if result := brokerCreator.LogIt(types.WARN, exampleStruct); !result {
		t.Error("Log not recorded")
	}
	if result := brokerCreator.LogIt(types.INFO, exampleStruct); !result {
		t.Error("Log not recorded")
	}
	time.Sleep(time.Second * 2)

	brokerCreator.Shutdown()
}

// TestBrokerCreatorWithJson tests logging JSON-encoded data with the BrokerCreator at the DEBUG and TRACE levels.
//
// It initializes a BrokerCreator with specified settings, creates an example struct,
// converts the struct to JSON, logs the JSON data at the DEBUG and TRACE levels,
// and checks if the log entries are recorded successfully. The test includes a sleep period to allow time for message processing.
func TestBrokerCreatorWithJson(t *testing.T) {
	brokerCreator, err := creators.NewBrokerCreator(brokers, "test", "Broker", 2, nil)
	if err != nil {
		t.Error(err)
	}

	exampleStruct := &struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Example Name",
		Age:  25,
	}

	jsonData, _ := json.Marshal(exampleStruct)

	if result := brokerCreator.LogIt(types.DEBUG, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
	if result := brokerCreator.LogIt(types.TRACE, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
	time.Sleep(time.Second * 2)
	brokerCreator.Shutdown()
}
