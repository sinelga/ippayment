package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/mitchellh/mapstructure"
	//	"time"
	//	"strconv"
	"domains"
	//	"sort"
)

var phonenumflag *string = flag.String("phonenum", "", "phonenume format 358...")

type ByAge []domains.Hit

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Created > a[j].Created }

func main() {
	flag.Parse() // Scan the arguments list

	phonenum := *phonenumflag

	if phonenum != "" {

		fmt.Println(phonenum)
		dir := "tiedotDB"

		tdDB, err := db.OpenDB(dir)
		if err != nil {
			panic(err)
		}

		mobclient := tdDB.Use("mobclients")

		for path := range mobclient.SecIndexes {
			fmt.Printf("I have an index on path %s\n", path)
		}

		var query interface{}
		var readBack interface{}
		queryStr := `{"eq": "` + phonenum + `","in": ["ClPhonenum"]}`

		json.Unmarshal([]byte(queryStr), &query)
		queryResult := make(map[uint64]struct{})
		if err := db.EvalQuery(query, mobclient, &queryResult); err != nil {
			panic(err)
		}
		var mobclientobj domains.MobClient
		for id := range queryResult {
			mobclient.Read(id, &readBack)
			fmt.Println(readBack)

			mobclientval := readBack.(map[string]interface{})

			err := mapstructure.Decode(mobclientval, &mobclientobj)
			if err != nil {
				panic(err)
			}
			fmt.Println("Hits ", mobclientobj.ClHits)
			if mobclientobj.ClResource == "mobilephone" {
				fmt.Println("Sms ", mobclientobj.ClSmsOut)

			}

		}

	}

}
