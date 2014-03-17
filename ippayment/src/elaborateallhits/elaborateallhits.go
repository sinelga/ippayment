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

		if quan_hits > 1 {

			for i := 0; i < quan_hits; i++ {

				println(i)

				var hit domains.Hits
//				type hittype domains.Hits

				bhit, _ := redis.Bytes(c.Do("RPOP", "hits"))

				err := json.Unmarshal(bhit, &hit)
				if err != nil {

					golog.Crit(err.Error())

				}

				if !checkcolexist.ChecExist(collections, hit.Msisdn) {

					println("create ", hit.Msisdn)
					if err := tdDB.Create(hit.Msisdn, 2); err != nil {
						//							panic(err)
						golog.Crit(err.Error())
					} else {

						collections = append(collections, hit.Msisdn)

					}

				} else {

					println("Update ", hit.Msisdn)
					msisdn := tdDB.Use(hit.Msisdn)
					docID, err := msisdn.Insert(map[string]interface{}{
						"Created": hit.Created,
						"Id": hit.Id,
						"Site": hit.Site,
						"Themes": hit.Themes,
						"Resource": hit.Resource,
					
					})

					if err != nil {
						panic(err)
					} else {
						
						println(docID)
					
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
