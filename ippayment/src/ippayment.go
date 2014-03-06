package main 

import (
    "net/http"
    "fmt"
    "jsonresponse"
//    "net/url"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
	msisdn := r.Header.Get("X-UP-CALLING-LINE-ID")
	callback := r.URL.Query().Get("callback")
	fmt.Fprintf(w, "callback ",callback)
    r.Header.Set("Content-Type", "application/json")
    jsonstr := jsonresponse.Response{"success": true, "msisdn": msisdn}
    fmt.Fprint(w, callback+"("+jsonstr.String()+")")
    
}

func main() {
    http.HandleFunc("/sonera", defaultHandler)
    http.ListenAndServe("107.170.68.93:80", nil)
}

