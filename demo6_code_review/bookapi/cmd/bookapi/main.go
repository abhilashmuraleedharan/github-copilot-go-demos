// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2023-10-05
package main

import (
    "log"
    "net/http"

    "bookapi/pkg/api"
)

func main() {
    router := api.NewRouter()
    serverAddr := ":8080"

    log.Printf("Starting server on %s...", serverAddr)
    if err := http.ListenAndServe(serverAddr, router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}