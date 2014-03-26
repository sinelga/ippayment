package main

import (
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"net/url"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
//var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var quantFlag = flag.Int("quant", 1, "quant can be > 1")

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	flag.Parse() // Scan the arguments list

	quant := *quantFlag

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}

	if quant_smsout, err := redis.Int(c.Do("LLEN", "smsout")); err != nil {

		golog.Crit(err.Error())

	} else {

		fmt.Println("quant_smsout", quant_smsout)

		for i := 0; i < quant; i++ {

			var smsout domains.SmsOut

			bsmsout, _ := redis.Bytes(c.Do("LPOP", "smsout"))

			err := json.Unmarshal(bsmsout, &smsout)
			var urlstr string
			if err != nil {

				golog.Crit(err.Error())

			} else {

				fmt.Println(smsout.Msisdn)

				if Url, err := url.Parse("http://79.125.27.200:9000"); err != nil {

					panic("boom")
				} else {
					Url.Path += "/"
					parameters := url.Values{}
					parameters.Add("msisdn", smsout.Msisdn)
					parameters.Add("text", smsout.Text)
					parameters.Add("from", smsout.From)
//					parameters.Add("site", smsout.)
					parameters.Add("provider", smsout.Provider)

					Url.RawQuery = parameters.Encode()
					urlstr = Url.String()

				}
				fmt.Println(urlstr)

//				resp, err := http.Get("http://79.125.27.200:9000/?msisdn=" + smsout.Msisdn)
				resp, err := http.Get(urlstr)
				defer resp.Body.Close()
				if err != nil {

					golog.Crit(err.Error())
				} else {
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						golog.Crit(err.Error())
					} else {

						golog.Info(string(body))

					}

				}

			}

		}

	}

}
