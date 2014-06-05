package main

import (
	"flag"
	"fmt"
	//	"github.com/HouzuoGuo/tiedot/db"
	"CreateHtmlGz"
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"hitshandler"
	"log"
	"log/syslog"
	//	"elaborateallhits"
	//	"time"
	//	"math/rand"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}

	//	elaborateallhits.ElabAllHits(*golog,c,*tdDB)
	//	tdDB.Flush()

	if quan_hits, err := redis.Int(c.Do("LLEN", "hits")); err != nil {

		golog.Crit(err.Error())

	} else {

		if quan_hits > 0 {

			for i := 0; i < quan_hits; i++ {

				var hit domains.Hit

				bhit, _ := redis.Bytes(c.Do("RPOP", "hits"))

				err := json.Unmarshal(bhit, &hit)
				if err != nil {

					golog.Crit(err.Error())

				} else {

					MobClientHtml := hitshandler.ElaborateHit(*golog, hit)
					CreateHtmlGz.CreateFile(*golog, MobClientHtml)

				}

			}

		}

	}

}
