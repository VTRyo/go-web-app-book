package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(r.body)
		if err != nil {
			log.Printf("Failed to log request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
