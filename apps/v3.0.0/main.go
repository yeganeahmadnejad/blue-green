package main

import (
    "fmt"
    "log"
	"time"
    "net/http"
)


func main() {

	healthy := false

	time.AfterFunc(time.Second*10, func() {
		healthy = true
	})


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my awesome app v3.0.0",)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
		if healthy {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	})
	
log.Fatal(http.ListenAndServe(":3000", nil))

}
