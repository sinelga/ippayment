package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"jsonresponse"
	"loadsubnets"
	"log"
	"log/syslog"
	"memdb"
	"net"
	"net/http"
	"net/http/fcgi"
	"sync"
	"domains"
)

var startOnce sync.Once
var ip string
var providersubnetarr []domains.ProviderSubnet

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	startOnce.Do(func() {
		startones(*golog)
	})

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}

	ipstr := req.RemoteAddr
	golog.Info("ip :"+ipstr)
	
	ipo := net.ParseIP(ipstr)
	
	for _,providersubnet := range providersubnetarr {
	
		if providersubnet.IpNet.Contains(ipo) {
		
			golog.Info("provider :"+providersubnet.Provider)
		
		}
	
	}
	
	

	msisdn := req.Header.Get("X-UP-CALLING-LINE-ID")
	golog.Info("msisdn " + msisdn)

	callback := req.URL.Query().Get("callback")
	site := req.URL.Query().Get("site")
	id := req.URL.Query().Get("id")
	resource := req.URL.Query().Get("resource")
	themes := req.URL.Query().Get("themes")
	provider := req.URL.Query().Get("provider")

	golog.Info("id: " + id + " site: " + site + " resource: " + resource + " themes: " + themes)

	req.Header.Set("Content-Type", "application/json")

	if callback != "" && msisdn != "" && provider != "" {

		jsonstrtrue := jsonresponse.Response{"success": true, "msisdn": msisdn}
		fmt.Fprint(resp, callback+"("+jsonstrtrue.String()+");")

		record := []string{id, msisdn, site, themes, resource, provider}
		memdb.InsertHit(*golog, c, record)

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

func startones(golog syslog.Writer) {

	golog.Info("startones: Start")
	fieldsarr := loadsubnets.Load(golog, "allsubnet.csv")

	for _, field := range fieldsarr {

		_, ipnet, _ := net.ParseCIDR(field[0])
		
		providersubnet := domains.ProviderSubnet {
		
			IpNet: *ipnet,
			Provider: field[1], 
		
		}
		
		providersubnetarr = append(providersubnetarr,providersubnet)
								
	
	}

}
