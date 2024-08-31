package transport

import (
	"encoding/json"
	"net/http"
)

type CommonResponse struct {
	Code    string      `json:"code,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SendResponseOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := CommonResponse{
		Code:    "00",
		Data:    data,
		Message: "OK",
	}

	encodeJsonErr := json.NewEncoder(w).Encode(resp)
	if encodeJsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SendResponseError(w http.ResponseWriter, err error) {
	resp := CommonResponse{
		Code:  "01",
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	encodeJsonErr := json.NewEncoder(w).Encode(resp)
	if encodeJsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
