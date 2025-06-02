package main

import (
    "my-go-project/internal/app"
    "log"
)

func main() {
    a := app.NewApp()
    if err := a.Start(); err != nil {
        log.Fatalf("Failed to start the application: %v", err)
    }
    defer a.Stop()
}