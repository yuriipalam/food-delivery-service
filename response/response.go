package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Basic struct {
	Message string `json:"message"`
}

func SendOK(w http.ResponseWriter, data any) {
	sendJSON(w, http.StatusOK, data)
}

func sendJSON(w http.ResponseWriter, code int, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		SendInternalServerError(w, errors.New("error with sending json"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
