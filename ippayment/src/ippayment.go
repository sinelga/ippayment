package main

import (
	"fmt"
	"jsonresponse"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"github.com/garyburd/redigo/redis"
	"memdb"
)


type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
	
		golog.Crit(err.Error())
		
	}

	msisdn := req.Header.Get("X-UP-CALLING-LINE-ID")
	golog.Info("msisdn " + msisdn)
	
	callback := req.URL.Query().Get("callback")
	site := req.URL.Query().Get("site")
	id := req.URL.Query().Get("id")
	resource := req.URL.Query().Get("resource")
	themes := req.URL.Query().Get("themes")
	provider :=req.URL.Query().Get("provider")
	
//	record := []string{id,site,themes,resource}
	
	golog.Info("id: "+id+" site: "+site+" resource: "+resource+" themes: "+themes)
		
	req.Header.Set("Content-Type", "application/json")
	
	if callback != ""  && msisdn != "" && provider!="" {
		
		jsonstrtrue := jsonresponse.Response{"success": true, "msisdn": msisdn}
		fmt.Fprint(resp, callback+"("+jsonstrtrue.String()+");")
		
		record := []string{id,msisdn,site,themes,resource,provider}
		memdb.InsertHit(*golog,c,record)
		
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
