package main 

import (
    "net/http"
    "fmt"
    "jsonresponse"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
	msisdn := r.Header.Get("X-UP-CALLING-LINE-ID")
//    fmt.Fprintf(w, r.Header.Get("X-UP-CALLING-LINE-ID"))
    r.Header.Set("Content-Type", "application/json")
    fmt.Fprint(w, jsonresponse.Response{"success": true, "msisdn": msisdn})
    
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe("107.170.68.93:80", nil)
}

