package utils

import(

	"encoding/json"
	"log"
	"net/http"
	"bytes"
	"io"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"

)

func WriteJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}

func writeJSONError(w http.ResponseWriter, v interface{}, code int) {
	w.WriteHeader(code)
	WriteJSON(w, v)
	return
}

func writeJSONErrorMessage(w http.ResponseWriter, message string, code int) {
	resp := &model.ErrorResponse{
		Error: &model.ErrorResponseMessage{
			Message: message,
		},
	}
	writeJSONError(w, resp, code)
}