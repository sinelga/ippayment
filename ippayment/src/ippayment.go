package main

import (
	"CreateHtmlGz"
	"domains"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"htmlhandler"
	"jsonresponse"
	"loadsubnets"
	"log"
	"log/syslog"
	"memdb"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
	"sync"
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

	provider := "NotMobile"

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}

	ipstr := strings.Split(req.RemoteAddr, ":")

	golog.Info("req.RequestURI " + req.URL.RequestURI())

	ipo := net.ParseIP(ipstr[0])

	for _, providersubnet := range providersubnetarr {

		if providersubnet.IpNet.Contains(ipo) {

			provider = providersubnet.Provider

		}

	}

	msisdn := req.Header.Get("X-UP-CALLING-LINE-ID")

	callback := req.URL.Query().Get("callback")
	site := req.URL.Query().Get("site")
	id := req.URL.Query().Get("id")
	resource := req.URL.Query().Get("resource")
	themes := req.URL.Query().Get("themes")

	golog.Info("id: " + id + " site: " + site + " resource: " + resource + " themes: " + themes + " provider: " + provider + " msisdn: " + msisdn)

	if provider == "MobileSonera" && site != "" && id != "" && msisdn != "" && themes != "" && resource != "" {

		golog.Info("Ok insert record in redis for Sonera")
		record := []string{id, msisdn, site, themes, resource, provider}
		memdb.InsertHit(*golog, c, record)

	}

//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Access-Control-Allow-Origin", "*")
//	resp.Header().Set("Access-Control-Allow-Origin", "*")

	if callback != "" && id != "" {

		jsonstrtrue := jsonresponse.Response{"success": true, "msisdn": msisdn, "provider": provider}
		
		
		fmt.Fprint(resp, callback+"("+jsonstrtrue.String()+");")
		golog.Info("Set CORS")
		resp.Header().Set("Access-Control-Allow-Origin", "*")

	}

	if strings.HasPrefix(req.URL.RequestURI(), "/blocktel") {

		golog.Info("tel: " + req.URL.Query().Get("tel"))
		tel := req.URL.Query().Get("tel")

		if strings.HasPrefix(tel, "358") {
			htmlfile := string("/home/juno/git/ippayment/ippayment/www/tel/" + tel + ".html.gz")
			mobclienthtml := htmlhandler.ParseHtmlGzFile(*golog, htmlfile)
			mobclienthtml.ClBlock = "Yes"

			CreateHtmlGz.CreateFile(*golog, mobclienthtml)
		} else {

			golog.Err("!!! Check req "+req.URL.RequestURI())
		}

	}

	if strings.HasPrefix(req.URL.RequestURI(), "/tel") {

		fmt.Fprint(resp, "Don't exist "+req.URL.RequestURI())

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

		providersubnet := domains.ProviderSubnet{

			IpNet:    *ipnet,
			Provider: field[1],
		}

		providersubnetarr = append(providersubnetarr, providersubnet)

	}

}
