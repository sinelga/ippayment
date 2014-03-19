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

type ByAge []domains.Hits

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

		phonenumcol := tdDB.Use(phonenum)

		for path := range phonenumcol.SecIndexes {
			fmt.Printf("I have an index on path %s\n", path)
		}

		var query interface{}
		var readBack interface{}
		queryStr := `{"has": ["ColCreated"],"limit":1000}`
		json.Unmarshal([]byte(queryStr), &query)
		queryResult := make(map[uint64]struct{})
		if err := db.EvalQuery(query, phonenumcol, &queryResult); err != nil {
			panic(err)
		}
		for id := range queryResult {
			phonenumcol.Read(id, &readBack)
			fmt.Println(readBack)

		}

		//		queryStr := `["all"]`
		//		queryStr := `{"has": ["Created","Id","Site","Themes","Resource"]}`
		//		queryStr := `{"has": ["Created","Id","Site","Themes","Resource"],"limit":1}`
		queryStr = `{"has": ["Created"],"limit":1000}`

		json.Unmarshal([]byte(queryStr), &query)
		queryResult = make(map[uint64]struct{})
		if err := db.EvalQuery(query, phonenumcol, &queryResult); err != nil {
			panic(err)
		}

		var hitsarr []domains.Hits

		for id := range queryResult {
			//
			var hit domains.Hits
			phonenumcol.Read(id, &readBack)

			vals := readBack.(map[string]interface{})

//			var result domains.Hits

			err := mapstructure.Decode(vals, &hit)
			if err != nil {
				panic(err)
			}
//			fmt.Println( result.Created)
//			fmt.Printf("%#v", result)
			//			var md domains.Metadata
			//
			//			var result domains.Hits
			//
			//			config := domains.DecoderConfig{
			//				Metadata: &md,
			//				Result:   &result,
			//			}
			//
			//			decoder, err := mapstructure.NewDecoder(&config)
			//			if err != nil {
			//				panic(err)
			//			}

//			for k, v := range vals {
//
//				switch vv := v.(type) {
//				case string:
//					fmt.Println(k, "is string", vv)
//				case int:
//					fmt.Println(k, "is int", vv)
//				case []interface{}:
//					fmt.Println(k, "is an array:")
//					for i, u := range vv {
//						fmt.Println(i, u)
//					}
//				default:
//					fmt.Println(k, "is of a type I don't know how to handle")
//				}
//
//			}

//			created := vals["Created"].(float64)
//			createdint64 := int64(created)
//			id := vals["Id"].(string)
//			site := vals["Site"].(string)
//			themes := vals["Themes"].(string)
//			resource := vals["Resource"].(string)

//			hit.Created = createdint64
//			hit.Id = id
//			hit.Site = site
//			hit.Themes = themes
//			hit.Resource = resource

			hitsarr = append(hitsarr, hit)

		}

		sort.Sort(ByAge(hitsarr))

		for _, hit := range hitsarr {

			createddate := time.Unix(hit.Created, 0)
			fmt.Println(createddate, hit.Id, hit.Site, hit.Themes, hit.Resource)

		}

	}

}
