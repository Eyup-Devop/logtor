package logtor_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Eyup-Devop/logtor"
	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

func TestGetLogCreatorListHandlerFunc(t *testing.T) {
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

	// Create a request to pass to your handler function
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.GetLogCreatorList(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseList []string
	err = json.NewDecoder(rw.Body).Decode(&responseList)
	if err != nil {
		t.Errorf("handler returned not json data")
	}
}

func TestGetCurrentLogCreatorHandlerFunc(t *testing.T) {
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

	// Create a request to pass to your handler function
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.GetCurrentLogCreator(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"current_log_creator":"Console"}`
	if rw.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rw.Body.String(), expected)
	}
}

func TestChangeActiveLogCreatorHandlerFunc(t *testing.T) {
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

	payload := map[string]string{
		"log_creator": "File",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Error(err.Error())
		return
	}

	// Create a request to pass to your handler function
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.ChangeActiveLogCreator(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"old_log_creator":"Console","current_log_creator":"File"}`
	if rw.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rw.Body.String(), expected)
	}
}

func TestGetLogLevelListHandlerFunc(t *testing.T) {
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

	// Create a request to pass to your handler function
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.GetLogLevelList(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseList []string
	err = json.NewDecoder(rw.Body).Decode(&responseList)
	if err != nil {
		t.Errorf("handler returned not json data")
	}
}

func TestGetActiveLogLevelHandlerFunc(t *testing.T) {
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

	// Create a request to pass to your handler function
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.GetActiveLogLevel(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"log_level":"TRACE"}`
	if rw.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rw.Body.String(), expected)
	}
}

func TestSetLogLevelHandlerFunc(t *testing.T) {
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

	logLevelPayload := "ERROR"

	// Create a request to pass to your handler function
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(logLevelPayload)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rw := httptest.NewRecorder()

	// Call your handler function, passing in the ResponseRecorder and the Request
	newLogtor.SetLogLevelHandlerFunc(rw, req)

	// Check the status code
	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"old_log_level":"TRACE","current_log_level":"ERROR"}`
	if rw.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rw.Body.String(), expected)
	}
}
