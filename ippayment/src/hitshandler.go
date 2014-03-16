package main

import (
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"elaborateallhits"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	dir := "tiedotDB"
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	var collection []string

	tdDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}
	for name := range tdDB.StrCol {
		fmt.Printf("I have a collection called %s\n", name)
		collection = append(collection,name)		
	}
	
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}
	
	elaborateallhits.ElabAllHits(*golog,c,collection)
	

}
