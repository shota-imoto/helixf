package supports

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(503)
	response_struct := ErrorResponse{Message: err.Error()}
	response_json, err := json.Marshal(response_struct)
	if err != nil {
		panic(err)
	}
	w.Write(response_json)
}

func UnauthorizedHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(401)

	response_struct := ErrorResponse{Message: err.Error()}
	response_json, err := json.Marshal(response_struct)
	if err != nil {
		panic(err)
	}
	w.Write(response_json)
}
