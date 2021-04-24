package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my awesome app v4.0.0",)
    })

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
        w.WriteHeader(500)
    })

    log.Fatal(http.ListenAndServe(":3000", nil))

}
