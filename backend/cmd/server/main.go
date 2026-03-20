
package main

import (
    "log"
    "net/http"
    "web-analyzer/internal/handler"
)

func main() {
    http.HandleFunc("/api/analyze", handler.AnalyzeHandler)
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
