package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HELLO FROM GO")
}

func main() {
    
    http.HandleFunc("/", handler)
    fmt.Println("Running Go Server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}