package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func messageHandler(message string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
    }
}

func main() {
    port := os.Getenv("PORT");
    msg := os.Getenv("MSG");

    fmt.Println("PORT:", port)
    fmt.Println("MSG:", msg)

    log.Fatal(http.ListenAndServe(":8080", messageHandler(msg)))
}
