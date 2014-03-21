package elaborateallhits

import (
	//"log"
	"checkcolexist"
	//	"checkindex"
	"domains"
	"encoding/json"
	"fmt"
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

		col := tdDB.Use("mobclients")

		if quan_hits > 0 {

			for i := 0; i < quan_hits; i++ {

				var hit domains.Hit

				bhit, _ := redis.Bytes(c.Do("RPOP", "hits"))

				err := json.Unmarshal(bhit, &hit)
				if err != nil {

					golog.Crit(err.Error())

				}

				docID := checkcolexist.ChecExist(col, hit.Msisdn)
				if docID == 0 {

					println("create ", hit.Msisdn)

					var hitsarr []domains.Hit
					var smsoutarr []domains.SmsOut

					hitsarr = append(hitsarr, hit)

					if hit.Resource == "mobilephone" {

						smsout := domains.SmsOut{
							SmsCreated: nowunix,
							Msisdn:     hit.Msisdn,
							From:       "070095943",
							Text:       "Soita! Miia. puh. 070095943",
							Provider:	hit.Provider,
						}
						pushsmsout.PushOut(golog, c, smsout)
						smsoutarr = append(smsoutarr, smsout)

						col.Insert(map[string]interface{}{
							"ClPhonenum": hit.Msisdn,
							"ClCreated":  hit.Created,
							"ClUpdated":  hit.Created,
							"ClThemes":   hit.Themes,
							"ClResource": hit.Resource,
							"ClProvider": hit.Provider,
							"ClBlock":    0,
							"ClHits":     hitsarr,
							"ClSmsOut":   smsoutarr,
						})

					} else {

						col.Insert(map[string]interface{}{
							"ClPhonenum": hit.Msisdn,
							"ClCreated":  hit.Created,
							"ClUpdated":  hit.Created,
							"ClThemes":   hit.Themes,
							"ClResource": hit.Resource,
							"ClProvider": hit.Provider,
							"ClBlock":    0,
							"ClHits":     hitsarr,
						})

					}

				} else {

					println("Update ", hit.Msisdn)

					var readBack interface{}

					col.Read(docID, &readBack)

					mobclientval := readBack.(map[string]interface{})
					var mobclientobj domains.MobClient

					err := mapstructure.Decode(mobclientval, &mobclientobj)
					if err != nil {
						panic(err)
					}

					fmt.Println("DECODE 2", mobclientobj.ClHits)

					mobclientobj.ClHits = append(mobclientobj.ClHits, hit)

					if hit.Resource == "mobilephone" {

						col.Update(docID, map[string]interface{}{
							"ClPhonenum": mobclientobj.ClPhonenum,
							"ClCreated":  mobclientobj.ClCreated,
							"ClUpdated":  nowunix,
							"ClThemes":   mobclientobj.ClThemes,
							"ClResource": mobclientobj.ClResource,
							"ClProvider": mobclientobj.ClProvider,
							"ClBlock":    mobclientobj.ClBlock,
							"ClHits":     mobclientobj.ClHits,
							"ClSmsOut":   mobclientobj.ClSmsOut,
						})

					} else {

						col.Update(docID, map[string]interface{}{
							"ClPhonenum": mobclientobj.ClPhonenum,
							"ClCreated":  mobclientobj.ClCreated,
							"ClUpdated":  nowunix,
							"ClThemes":   mobclientobj.ClThemes,
							"ClResource": mobclientobj.ClResource,
							"ClProvider": mobclientobj.ClProvider,
							"ClBlock":    mobclientobj.ClBlock,
							"ClHits":     mobclientobj.ClHits,
						})

					}

				}

			}

		}

	}

}
