package nomockutil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteObjectOrErr(w http.ResponseWriter, dataObj interface{}, errStatusCode int, err error) {
	if err != nil {
		WriteError(errStatusCode, w, err)
	} else {
		writeObject(w, dataObj)
	}
}

func writeObject(w http.ResponseWriter, responseObj interface{}) {
	responseBytes, err := json.Marshal(responseObj)
	if err != nil {
		WriteError(http.StatusInternalServerError, w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", responseBytes)
	}
}

func WriteError(statusCode int, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	errMap := make(map[string]string)
	errMap["error"] = err.Error()
	data, marshalErr := json.Marshal(errMap)
	if marshalErr == nil {
		w.Write(data)
	} else {
		w.Write([]byte("{}"))
	}
}
