package logtor

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Eyup-Devop/logtor/types"
)

func (l *Logtor) GetLogCreatorList(w http.ResponseWriter, r *http.Request) {
	result := []string{}
	l.changeMutex.RLock()
	defer l.changeMutex.RUnlock()
	for k := range l.logCreatorList {
		result = append(result, string(k))
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func (l *Logtor) GetCurrentLogCreator(w http.ResponseWriter, r *http.Request) {
	l.changeMutex.RLock()
	defer l.changeMutex.RUnlock()
	if len(l.logCreatorList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := struct {
		CurrentLogCreator string `json:"current_log_creator"`
	}{
		CurrentLogCreator: string(l.activeLogCreator.LogName()),
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func (l *Logtor) ChangeActiveLogCreator(w http.ResponseWriter, r *http.Request) {
	l.changeMutex.RLock()
	defer l.changeMutex.RUnlock()
	if len(l.logCreatorList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || len(payload) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	oldLogCreator := string(l.activeLogCreator.LogName())
	var currentLogCreator string
	if v, ok := payload["log_creator"]; ok {
		l.changeMutex.RUnlock()
		if l.ChangeLogCreator(types.LogCreatorName(v)) {
			currentLogCreator = v
		} else {
			currentLogCreator = oldLogCreator
		}
		l.changeMutex.RLock()
	}

	result := struct {
		OldLogCreator     string `json:"old_log_creator"`
		CurrentLogCreator string `json:"current_log_creator"`
	}{
		OldLogCreator:     oldLogCreator,
		CurrentLogCreator: currentLogCreator,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func (l *Logtor) GetLogLevelList(w http.ResponseWriter, r *http.Request) {
	jsonResult, err := json.Marshal(types.LogLevelList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

func (l *Logtor) GetActiveLogLevel(w http.ResponseWriter, r *http.Request) {
	l.changeMutex.RLock()
	defer l.changeMutex.RUnlock()
	if len(l.logCreatorList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := struct {
		LogLevel string `json:"log_level"`
	}{
		LogLevel: string(l.LogLevel()),
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func (l *Logtor) SetLogLevelHandlerFunc(w http.ResponseWriter, r *http.Request) {
	l.changeMutex.RLock()
	if len(l.logCreatorList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var oldLogLevel string = string(l.LogLevel())

	l.changeMutex.RUnlock()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytePayload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payload := string(bytePayload)
	var currentLogLevel string
	if l.SetLogLevel(types.LogLevel(payload)) {
		currentLogLevel = payload
	} else {
		currentLogLevel = oldLogLevel
	}

	result := struct {
		OldLogLevel     string `json:"old_log_level"`
		CurrentLogLevel string `json:"current_log_level"`
	}{
		OldLogLevel:     oldLogLevel,
		CurrentLogLevel: currentLogLevel,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}
