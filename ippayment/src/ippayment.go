package main

import (
	"fmt"
	"jsonresponse"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
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
	site := req.URL.Query().Get("site")
	id := req.URL.Query().Get("id")
	resource := req.URL.Query().Get("resource")
	themes := req.URL.Query().Get("themes")
	
	golog.Info("id: "+id+" site: "+site+" resource: "+resource+" themes: "+themes)
		
	req.Header.Set("Content-Type", "application/json")
	
	if callback != ""  && msisdn != "" {
		
		jsonstrtrue := jsonresponse.Response{"success": true, "msisdn": msisdn}
		fmt.Fprint(resp, callback+"("+jsonstrtrue.String()+");")
	} else {
		jsonstrfalse := jsonresponse.Response{"success": false, "msisdn": ""}
		fmt.Fprint(resp, callback+"("+jsonstrfalse.String()+");")		
	}

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
