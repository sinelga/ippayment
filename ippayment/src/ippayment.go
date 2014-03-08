package main

import (
	"fmt"
	"jsonresponse"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	//    "net/url"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	msisdn := req.Header.Get("X-UP-CALLING-LINE-ID")

	golog.Info("msisdn " + msisdn)

	callback := req.URL.Query().Get("callback")
	//	fmt.Fprintf(w, "callback ",r.URL.RawQuery)
	req.Header.Set("Content-Type", "application/json")
	//	var jsonstr map[string]interface{}
	if callback != "" {
		jsonstr := jsonresponse.Response{"success": true, "msisdn": msisdn}
		fmt.Fprint(resp, callback+"("+jsonstr.String()+");")
	} else {
		http.Error(resp, "Too Many Requests", 429)
	}

}

// Default Request Handler
//func defaultHandler(w http.ResponseWriter, r *http.Request) {
//	//    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
//	msisdn := r.Header.Get("X-UP-CALLING-LINE-ID")
//	callback := r.URL.Query().Get("callback")
//	//	fmt.Fprintf(w, "callback ",r.URL.RawQuery)
//	r.Header.Set("Content-Type", "application/json")
//	jsonstr := jsonresponse.Response{"success": true, "msisdn": msisdn}
//	fmt.Fprint(w, callback+"("+jsonstr.String()+");")
//
//}

func main() {
	//    http.HandleFunc("/sonera", defaultHandler)
	//    http.ListenAndServe("107.170.68.93:80", nil)

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
