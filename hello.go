package main

import (
    "log"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "{\"Status\": \"OK\"}")
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
