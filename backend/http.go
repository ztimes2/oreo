package main

import (
	"net/http"
	"log"
)

const (
	headerContentType = "Content-Type"

	contentTypeJSON = "application/json"
)

func writeJSON(w http.ResponseWriter, statusCode int, resp interface{}) {
	if resp == nil {
		w.WriteHeader(statusCode)
		return
	}
	
	body, err := json.Marshal(resp)
	if err != nil {
		writeUnexpectedError(w, err)
		return
	}

	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)
	w.Write(body)
}

func writeUnexpectedError(w http.ResponseWriter, err error) {
	log.Printf("unexpected error: %v", err)

	body, _ := json.Marshal(errorResponse{
		Description: "Something went wrong. Please, try again later."
	})

	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(body)
}

type errorResponse struct {
	Description string `json:"err_description"`
}

func writeError(w http.ResponseWriter, statusCode int, e errorResponse) {
	writeJSON(w, statusCode, e)
}
