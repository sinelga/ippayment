package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/mitchellh/mapstructure"
	"time"
	//	"strconv"
	"domains"
	"sort"
)

var phonenumflag *string = flag.String("phonenum", "", "phonenume format 358...")

type ByAge []domains.Hit

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Created < a[j].Created }

type ByAgeSms []domains.SmsOut

func (a ByAgeSms) Len() int           { return len(a) }
func (a ByAgeSms) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAgeSms) Less(i, j int) bool { return a[i].SmsCreated < a[j].SmsCreated }

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

			mobclientval := readBack.(map[string]interface{})

			err := mapstructure.Decode(mobclientval, &mobclientobj)
			if err != nil {
				panic(err)
			}

			sort.Sort(ByAge(mobclientobj.ClHits))
			fmt.Println("Hits ")
			for _, hit := range mobclientobj.ClHits {

				fmt.Println(time.Unix(hit.Created, 0), hit.Site, hit.Resource)

			}

			if mobclientobj.ClResource == "mobilephone" {
//				fmt.Println("Sms ", mobclientobj.ClSmsOut)
				fmt.Println("SmsOut")
				sort.Sort(ByAgeSms(mobclientobj.ClSmsOut))
				for _, smsout := range mobclientobj.ClSmsOut {
					fmt.Println(time.Unix(smsout.SmsCreated, 0),smsout.From,smsout.Text)
				
				}
				
			}

		}

	}

}
