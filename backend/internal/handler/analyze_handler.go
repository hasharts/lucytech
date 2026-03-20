package handler

import (
    "encoding/json"
    "net/http"

    "web-analyzer/internal/service"
)

type Request struct {
    URL string `json:"url"`
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight request
    if r.Method == "OPTIONS" {
        return
    }

    var req Request

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    result, err := service.Analyze(req.URL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadGateway)
        return
    }

    json.NewEncoder(w).Encode(result)
}