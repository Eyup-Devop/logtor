package creators_test

import (
	"encoding/json"
	"testing"

	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

func TestFileRecorderWithString(t *testing.T) {
	fileRecorder, err := creators.NewFileCreator("./temp/temp.log", "File", 3, 5)
	if err != nil {
		t.Error(err)
	}
	if result := fileRecorder.LogIt(types.ERROR, "Example File Log Message"); !result {
		t.Error("Log not recorded")
	}
}

func TestFileRecorderWithStruct(t *testing.T) {
	fileRecorder, err := creators.NewFileCreator("./temp/temp.log", "File", 3, 5)
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

	if result := fileRecorder.LogIt(types.WARN, exampleStruct); !result {
		t.Error("Log not recorded")
	}
	if result := fileRecorder.LogIt(types.INFO, exampleStruct); !result {
		t.Error("Log not recorded")
	}
}

func TestFileRecorderWithJson(t *testing.T) {
	fileRecorder, err := creators.NewFileCreator("./temp/temp.log", "File", 3, 5)
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

	if result := fileRecorder.LogIt(types.DEBUG, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
	if result := fileRecorder.LogIt(types.TRACE, string(jsonData)); !result {
		t.Error("Log not recorded")
	}
}
