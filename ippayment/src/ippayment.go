package main 

import (
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
    fmt.Fprintf(w, r.Header.Get("X-UP-CALLING-LINE-ID"))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":80", nil)
}

