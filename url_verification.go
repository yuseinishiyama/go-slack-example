package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type urlVerification struct {
}

func (h urlVerification) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := make(map[string]interface{})
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("[ERROR] Failed to decode payload")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data["challenge"].(string)))
}
