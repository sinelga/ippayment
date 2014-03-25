package main

import (
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"elaborateallhits"
	"time"
	"math/rand"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)		
	}
	rand.Seed(time.Now().UTC().UnixNano())
	
	dir := "tiedotDB"
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	tdDB, err := db.OpenDB(dir)
	defer tdDB.Close()
	if err != nil {
		panic(err)
	}
	
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}
	
	elaborateallhits.ElabAllHits(*golog,c,*tdDB)
	tdDB.Flush()

	
}
