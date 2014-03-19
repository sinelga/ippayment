package elaborateallhits

import (
	//"log"
	"checkcolexist"
	"checkindex"
	"domains"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"github.com/mitchellh/mapstructure"
	"log/syslog"
	"pushsmsout"
	"time"
)

func ElabAllHits(golog syslog.Writer, c redis.Conn, tdDB db.DB, collections []string) {

	if quan_hits, err := redis.Int(c.Do("LLEN", "hits")); err != nil {

		golog.Crit(err.Error())

	} else {

		nowunix := time.Now().Unix()

		if quan_hits > 0 {

			for i := 0; i < quan_hits; i++ {

				var hit domains.Hits

				bhit, _ := redis.Bytes(c.Do("RPOP", "hits"))

				err := json.Unmarshal(bhit, &hit)
				if err != nil {

					golog.Crit(err.Error())

				}

				if !checkcolexist.ChecExist(collections, hit.Msisdn) {

					println("create ", hit.Msisdn)
					if err := tdDB.Create(hit.Msisdn, 2); err != nil {

						golog.Crit(err.Error())
					} else {

						collections = append(collections, hit.Msisdn)

						checkindex.CheckIndx(golog, tdDB, hit)

						msisdn := tdDB.Use(hit.Msisdn)
						//						if err := msisdn.Index([]string{"Created"}); err != nil {
						//							panic(err)
						//						}
						//						if err := msisdn.Index([]string{"Id"}); err != nil {
						//							panic(err)
						//						}
						//						if err := msisdn.Index([]string{"Site"}); err != nil {
						//							panic(err)
						//						}
						//						if err := msisdn.Index([]string{"Themes"}); err != nil {
						//							panic(err)
						//						}
						//						if err := msisdn.Index([]string{"Resource"}); err != nil {
						//							panic(err)
						//						}
						//						if err := msisdn.Index([]string{"ColCreated"}); err != nil {
						//							panic(err)
						//						}

						msisdn.Insert(map[string]interface{}{
							"ColCreated":  hit.Created,
							"ColUpdated":  hit.Created,
							"ColThemes":   hit.Themes,
							"ColResource": hit.Resource,
							"ColHits":     0,
						})

						msisdn.Insert(map[string]interface{}{
							"Created":  hit.Created,
							"Id":       hit.Id,
							"Site":     hit.Site,
							"Themes":   hit.Themes,
							"Resource": hit.Resource,
						})

						if hit.Resource == "mobilephone" {

							smsout := domains.SmsOut{
								SmsCreated: nowunix,
								Msisdn:     hit.Msisdn,
								From:       "070095943",
								Text:       "",
							}
							pushsmsout.PushOut(golog, c, smsout)

							msisdn.Insert(map[string]interface{}{
								"SmsCreated": smsout.SmsCreated,
								"Msisdn":     smsout.Msisdn,
								"From":       smsout.From,
								"Text":       smsout.Text,
							})

						}

					}

				} else {

					println("Update ", hit.Msisdn, hit.Created)

					msisdn := tdDB.Use(hit.Msisdn)
					_, err := msisdn.Insert(map[string]interface{}{
						"Created":  hit.Created,
						"Id":       hit.Id,
						"Site":     hit.Site,
						"Themes":   hit.Themes,
						"Resource": hit.Resource,
					})
					if err != nil {
						panic(err)
					}

					var query interface{}
					var readBack interface{}
					queryStr := `{"has": ["ColCreated"],"limit":1000}`

					json.Unmarshal([]byte(queryStr), &query)
					queryResult := make(map[uint64]struct{})
					if err := db.EvalQuery(query, msisdn, &queryResult); err != nil {

						golog.Crit(err.Error())
						//			tdDB.Drop(name)

					}

					for id := range queryResult {
						var collection domains.Collection
						msisdn.Read(id, &readBack)
						vals := readBack.(map[string]interface{})
						err := mapstructure.Decode(vals, &collection)
						if err != nil {
							panic(err)
						}
						println("update2 ", collection.ColCreated)

						err = msisdn.Update(id, map[string]interface{}{
							"ColCreated":  collection.ColCreated,
							"ColUpdated":  nowunix,
							"ColThemes":   collection.ColThemes,
							"ColResource": collection.ColResource,
							"ColHits":     collection.ColHits + 1,
						})
						if err != nil {
							panic(err)
						}

					}

				}

			}

		}

	}

}


