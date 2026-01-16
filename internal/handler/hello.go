package handler

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello from Go Web App Boilerplate!",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
