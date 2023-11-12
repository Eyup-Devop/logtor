// Package creators_test provides test cases for the logtor package, focusing on the BaseCreator implementation.
//
// It includes test functions for logging various types of messages using the BaseCreator,
// including string messages, structured data in a struct, and JSON-encoded data.
//
// Test Functions:
//   - TestBaseCreatorWithString: Tests logging a string message with the BaseCreator at the ERROR level.
//   - TestBaseCreatorWithStruct: Tests logging a struct with the BaseCreator at the WARN and INFO levels.
//   - TestBaseCreatorWithJson: Tests logging JSON-encoded data with the BaseCreator at the DEBUG and TRACE levels.
package creators_test

import (
	"encoding/json"
	"testing"

	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

// TestBaseCreatorWithString tests logging a string message with the BaseCreator at the ERROR level.
//
// It initializes a BaseCreator with specified settings, logs an example string message at the ERROR level,
// and checks if the log entry is recorded successfully.
func TestBaseCreatorWithString(t *testing.T) {
	baseCreator, err := creators.NewBaseCreator("Console", 2, 5)
	if err != nil {
		t.Error(err)
	}
	if result := baseCreator.LogIt(types.ERROR, "Example Log Message"); !result {
		t.Error("Log not recorded")
	}
}

// TestBaseCreatorWithStruct tests logging a struct with the BaseCreator at the WARN and INFO levels.
//
// It initializes a BaseCreator with specified settings, creates an example struct,
// logs the struct at the WARN and INFO levels, and checks if the log entries are recorded successfully.
func TestBaseCreatorWithStruct(t *testing.T) {
	baseCreator, err := creators.NewBaseCreator("Console", 2, 5)
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

	if result := baseCreator.LogIt(types.WARN, exampleStruct); !result {
		t.Error("Log not recorded")
	}
	if result := baseCreator.LogIt(types.INFO, exampleStruct); !result {
		t.Error("Log not recorded")
	}
}

// TestBaseCreatorWithJson tests logging JSON-encoded data with the BaseCreator at the DEBUG and TRACE levels.
//
// It initializes a BaseCreator with specified settings, creates an example struct,
// converts the struct to JSON, logs the JSON data at the DEBUG and TRACE levels,
// and checks if the log entries are recorded successfully.
func TestBaseCreatorWithJson(t *testing.T) {
	baseCreator, err := creators.NewBaseCreator("Console", 2, 5)
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

	if result := baseCreator.LogIt(types.DEBUG, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
	if result := baseCreator.LogIt(types.TRACE, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
}
