package elaborateallhits

import (
	//"log"
	"checkcolexist"
	"domains"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func ElabAllHits(golog syslog.Writer, c redis.Conn, tdDB db.DB, collections []string) {

	if quan_hits, err := redis.Int(c.Do("LLEN", "hits")); err != nil {

		golog.Crit(err.Error())

	} else {

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

						msisdn := tdDB.Use(hit.Msisdn)
						if err := msisdn.Index([]string{"Created"}); err != nil {
							panic(err)
						}
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
						if err := msisdn.Index([]string{"ColCreated"}); err != nil {
							panic(err)
						}
						

						msisdn.Insert(map[string]interface{}{
							"ColCreated":  hit.Created,
							"ColUpdated":  hit.Created,
							"ColThemes":   hit.Themes,
							"ColResource": hit.Resource,
							"ColHits":     0,
						})

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

				}

			}

		}

	}

}

func insertHit(golog syslog.Writer, tdDB db.DB, hit domains.Hits) {

	//		msisdn :=
	//		docID, err :=sites.Insert(map[string]hit)
	//
	//			if err != nil {
	//		panic(err)
	//	}

}
