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


type ByAge []domains.Collection

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].ColCreated > a[j].ColCreated }


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
	if err != nil {
		panic(err)
	}


	var allcollectionarr []domains.Collection

	for name := range tdDB.StrCol {
		fmt.Printf("I have a collection called %s\n", name)
		phonenumcol := tdDB.Use(name)
		var query interface{}
		var readBack interface{}
		queryStr := `{"has": ["ColCreated"],"limit":1000}`

		json.Unmarshal([]byte(queryStr), &query)
		queryResult := make(map[uint64]struct{})
		if err := db.EvalQuery(query, phonenumcol, &queryResult); err != nil {

			golog.Crit(err.Error())
			tdDB.Drop(name)

		}
		for id := range queryResult {

			var collection domains.Collection
			phonenumcol.Read(id, &readBack)
			vals := readBack.(map[string]interface{})
			err := mapstructure.Decode(vals, &collection)
			if err != nil {
				panic(err)
			}
			collection.ColPhonenum = name
			allcollectionarr = append(allcollectionarr,collection)
			
//			fmt.Println(collection.ColResource)

		}

	}
	sort.Sort(ByAge(allcollectionarr))
	
	for _,col :=range allcollectionarr {
		
		createddate := time.Unix(col.ColCreated, 0)
		updateddate := time.Unix(col.ColUpdated, 0)
		fmt.Println(createddate,col.ColPhonenum,col.ColResource,col.ColHits,updateddate )
	
	
	}
	
		

}
