package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func errorHandler(w http.ResponseWriter, code int, errorMessage string) {
	if code > 499 {
		fmt.Printf("Error marshalling response %v", errorMessage)
	}
	responseHandler(w, code, ErrorResponse{
		Error: errorMessage,
	})
}

func responseHandler(w http.ResponseWriter, code int, reponseBody interface{}) {
	jsonResponse, err := json.Marshal(reponseBody)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Error marshalling response %v", reponseBody)
		w.WriteHeader(500)
		w.Write([]byte("Error marshalling response"))
		return
	}
	w.WriteHeader(code)
	w.Write(jsonResponse)
}