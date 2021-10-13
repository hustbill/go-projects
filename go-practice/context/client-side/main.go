// context/client-side/main.go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func main() {
    rootCtx := context.Background()
    req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
    if err != nil {
        panic(err)
    }
    // create context
    ctx, cancel := context.WithTimeout(rootCtx, 50*time.Millisecond)
    defer cancel()
    // attach context to our request
    req = req.WithContext(ctx)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    fmt.Println("resp received", resp)
}