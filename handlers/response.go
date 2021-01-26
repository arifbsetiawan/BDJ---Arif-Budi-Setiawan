package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct{}

func (r *Response) ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (r *Response) ResponseError(w http.ResponseWriter, code int, message string) {
	codeString := strconv.Itoa(code)
	response := DefaultDTO{
		Code:    codeString,
		Message: message,
	}
	r.ResponseJSON(w, code, response)
}

type DefaultDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
