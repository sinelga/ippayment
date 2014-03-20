package main

import (
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/mitchellh/mapstructure"
	"log"
	"log/syslog"
	"sort"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")


type ByAge []domains.MobClient
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].ClCreated > a[j].ClCreated }


func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	dir := "tiedotDB"

	tdDB, err := db.OpenDB(dir)
	defer tdDB.Close()
	if err != nil {
		panic(err)
	}

	col := tdDB.Use("mobclients")
	queryStr := `["all"]`
	



	var allmobclientsarr []domains.MobClient

//	for name := range tdDB.StrCol {
//		fmt.Printf("I have a collection called %s\n", name)
//		phonenumcol := tdDB.Use(name)
		var query interface{}
		var readBack interface{}
//		queryStr := `{"has": ["ColCreated"],"limit":1000}`
//
		json.Unmarshal([]byte(queryStr), &query)
		queryResult := make(map[uint64]struct{})
		if err := db.EvalQuery(query, col, &queryResult); err != nil {

			golog.Crit(err.Error())
//			tdDB.Drop(name)

		}
		for id := range queryResult {
//
			var mobclientobj domains.MobClient
			col.Read(id, &readBack)
			vals := readBack.(map[string]interface{})
			err := mapstructure.Decode(vals, &mobclientobj)
			if err != nil {
				panic(err)
			}
//			collection.ColPhonenum = name
//			allcollectionarr = append(allcollectionarr,collection)
			allmobclientsarr = append(allmobclientsarr,mobclientobj) 
//
//
		}

//	}
	sort.Sort(ByAge(allmobclientsarr))
//	
	for _,mobclient :=range allmobclientsarr {
//		
		createddate := time.Unix(mobclient.ClCreated, 0)
		updateddate := time.Unix(mobclient.ClUpdated, 0)
		fmt.Println(createddate,mobclient.ClPhonenum,mobclient.ClResource,len(mobclient.ClHits),updateddate )
//	
//	
	}
	
		

}
